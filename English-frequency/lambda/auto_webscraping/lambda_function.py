import sys
sys.path.append("lib.bs4")

from bs4 import BeautifulSoup
import requests,datetime,os,boto3

class Media:
    ID: int
    Name: str
    RSSUrl: str
    SelectRule: str

    def __init__(self,i,n,r,s):
        self.ID = i
        self.Name = n
        self.RSSUrl = r
        self.SelectRule = s

    

def GetArticleUrlFromRSS(rss_url):
    resultLinks = []
    req = requests.get(rss_url)
    selectRule = 'link'
    if req.status_code != 200:
        print(f"GetArticleUrlFromRSS ERROR: Request Failed. Status Code is {req.status_code}")
        return resultLinks
    
    soup = BeautifulSoup(req.text,"xml")
    links = soup.select(selectRule)

    for l in links:
        resultLinks.append(l.get_text())

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

def OutputStringListToS3(article_list,key):
    try:
        for l in article_list:
            obj = s3.Object(BUCKET,key)
            obj.put(Body=l)
    except Exception as e:
        print(f"OutputStringListToS3 ERROR: {e}")
        return -1



s3 = boto3.resource('s3')
MEDIALIST = []
MEDIALIST.append(Media(1,"BSS","http://feeds.bbci.co.uk/news/rss.xml#",'body > div > div > div > div > div > div > div > main > article > div > div > p[class*="Paragraph"]'))
BUCKET = 'nii-dev' 

def lambda_handler(event, context):

    media = MEDIALIST[0]
    urls = GetArticleUrlFromRSS(media.RSSUrl)
    fileNum = 1
    keypath = "./english-frequency/articles/" + media.Name + "/" + datetime.datetime.now().strftime('%Y-%m-%d')
    # os.makedirs(dirpath, exist_ok=True)
    for url in urls:
        articleList = GetArticle(url,media.SelectRule)
        filename = keypath+  "/" + str(fileNum)
        result = OutputStringListToS3(articleList,filename)
        if result != 0:
            break
        fileNum += 1
    return




