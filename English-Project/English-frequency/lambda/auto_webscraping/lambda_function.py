import sys
sys.path.append("lib.bs4")

from bs4 import BeautifulSoup
from botocore.config import Config
import requests,datetime,os,boto3,xmltodict

class Media:
    Provider: str
    RSSUrl: str
    SelectRule: str

    def __init__(self,i,n,r,s):
        self.ID = i
        self.Name = n
        self.RSSUrl = r
        self.SelectRule = s

# Media情報
MEDIALIST = []
MEDIALIST.append(Media(1,"BSS","http://feeds.bbci.co.uk/news/rss.xml#",'body > div > div > div > div > div > div > div > main > article > div > div > p[class*="Paragraph"]'))


    

def GetArticleUrlFromRSS(rss_url):
    resultLinks = []
    req = requests.get(rss_url)
    selectRule = 'link'

    if req.status_code != 200:
        print(f"GetArticleUrlFromRSS ERROR: Request Failed. Status Code is {req.status_code}")
        return resultLinks
    
    dict_xml_data = xmltodict.parse(req.text)
    items = dict_xml_data["rss"]["channel"]["item"]

    for i in items:
        resultLinks.append(i["link"])

    return resultLinks


def GetArticle(url,rule):
    resultArticles = []
    
    req = requests.get(url)

    if req.status_code != 200:
        print(f"GetArticle ERROR: Request Failed. Status Code is {req.status_code}")


    soup = BeautifulSoup(req.text,"html.parser")
    article = soup.select(rule)

    for a in article:
        resultArticles.append(a.get_text())

    return resultArticles


def OutputStringListToFile(article_list,file_name):
    try:
        with open(file_name,mode='w') as f:
            for l in article_list:
                f.write(l)
        return 0
    except Exception as e:
        print(f"OutputStringListToFile ERROR: {e}")
        return -1

def OutputStringListToS3(article_list,bucket,key,s3):
    try:
        article_body = " ".join(article_list)
        obj = s3.Object(bucket,key)
        obj.put(Body=article_body)
        return 0
        
    except Exception as e:
        print(f"OutputStringListToS3 ERROR: {e}")
        return -1




def lambda_handler(event, context):
    s3 = boto3.resource('s3', config=Config(signature_version='s3v4'))
    bucket = event["bucket"]
    provider = event["provider"]


    print("INFO: Bucket:" + bucket)
    print("INFO: Provider:" + provider)

    media = -1

    for m in MEDIALIST:
        if m.Name == provider:
            media = m

    if media == -1:
        print("ERROR: Provider Name doesn't exist")
        sys.exit(-1)

    urls = GetArticleUrlFromRSS(media.RSSUrl)
    fileNum = 1
    keypath = "english-frequency/articles/" + media.Name + "/" + datetime.datetime.now().strftime('%Y-%m-%d')

    for url in urls:
        articleList = GetArticle(url,media.SelectRule)
        filename = keypath+  "/" + str(fileNum) + ".txt"
        result = OutputStringListToS3(articleList,bucket,filename,s3)
        if result != 0:
            break
        fileNum += 1

    return




