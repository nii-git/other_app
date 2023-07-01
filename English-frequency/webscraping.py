from bs4 import BeautifulSoup
import requests,re

def GetArticleUrl():
    None


def GetArticle(url,rule):
    resultArticle = []
    if rule == "BBS":
        selectRule = 'body > div > div > div > div > div > div > div > main > article > div > div > p[class*="Paragraph"]'
    else:
        print("GetArticle ERROR: Rule is invalid")
        return resultArticle
    
    req = requests.get(url)

    if req.status_code != 200:
        print(f"GetArticle ERROR: Request Failed. Status Code is {req.status_code}")


    soup = BeautifulSoup(req.text,"html.parser")
    article = soup.select(selectRule)

    for a in article:
        resultArticle.append(a.get_text())
        # print(a.get_text())

    return resultArticle

if __name__ == '__main__':
    article = GetArticle("https://www.bbc.com/news/world-europe-66070033?at_medium=RSS&at_campaign=KARANGA","BBS")
    print(article[0])