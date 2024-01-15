import pymysql
import sys
import boto3
import os
import datetime
import glob,re
import pandas as pd


# todo: 環境変数に移す
ENDPOINT="free-db.cfoyprdoczrs.ap-northeast-1.rds.amazonaws.com"
PORT="3306"
USER="admin"
REGION="ap-northeast"
DBNAME="english_frequency"


def init(event):
    provider = event["provider"]
    format = "%Y-%m-%dT%H:%M:%SZ"
    date_utc = datetime.datetime.strptime(event["datetime"],format)
    date_jst = date_utc + datetime.timedelta(hours=9)
    date = date_jst.strftime("%Y-%m-%d")
    return provider,date

def getpath(s3client,provider,day):
    resp = s3client.list_objects_v2(Bucket="nii-dev",Prefix="english-frequency/articles/"+provider+"/"+day+"/")
    files = []
    for c in resp["Contents"]:
        if c["Size"] <= 0:
            continue
        files.append(c["Key"])
    return files


def local_output(df_groupby,provider,day):
    with open("./result/" + provider + "/" + day + "/result.csv","a") as f:
        for i in range(0,df_groupby.count()):
            index = df_groupby.index[i]
            values = df_groupby.values[i]
            f.write(index + "," + str(values) + "\n")


def db_error_killed_process(conn,err_msg,rollback=True):
    print(err_msg)
    if rollback:
        conn.rollback()
    sys.exit(-1)

def output(df_groupby,provider,day):
    with pymysql.connect(host=ENDPOINT, user=USER, passwd=os.environ["db_pass"], db=DBNAME) as conn:
        with conn.cursor() as cur:
            # fixme:未分類のカテゴリIDを取得 
            result_count = cur.execute("SELECT word_type_id FROM mst_wordtype WHERE word_type_name = '未分類'")
            if result_count != 1:
                db_error_killed_process(conn,"ERROR: Can't load uncategory-id",rollback=False)
            uncategory_id = cur.fetchall()[0][0]

            # providerのidを確認
            result_count = cur.execute("SELECT id FROM mst_provider WHERE id = %s",(provider,))
            if result_count != 1:
                db_error_killed_process(conn,"ERROR: Can't load provider-id",rollback=False)


            for i in range(0,df_groupby.count()):
                word = df_groupby.index[i]
                count = df_groupby.values[i]
                result_count = cur.execute("SELECT id FROM word WHERE word = %s", (word,))

                if result_count == 0:
                    # wordテーブルにない場合
                    result_count = cur.execute("INSERT INTO word(word,word_type) VALUES (%s,%s)",(word,uncategory_id))
                    if result_count != 1:
                        db_error_killed_process(conn,"ERROR: Failed to write word table")
                    # 再度wordidを取得する
                    result_count = cur.execute("SELECT id FROM word WHERE word = %s", (word,))
                    if result_count != 1:
                        print(cur.fetchall())
                        db_error_killed_process(conn,"ERROR: Failed to get wordid")
                    word_id = cur.fetchall()[0][0]
                elif result_count > 1:
                    db_error_killed_process(conn,"ERROR: Inconsistency in word table")
                else:
                    word_id = cur.fetchall()[0][0]
                
                # frequencyテーブルに登録
                result_count = cur.execute("INSERT INTO frequency(provider_id,word_id,count,date) VALUES (%s,%s,%s,%s)",(provider,word_id,count,day))
                if result_count != 1:
                    db_error_killed_process("ERROR: Failed to insert frequency data")

            conn.commit()

                    
    return




def lambda_handler(event, context):
    provider,date = init(event)

    s3 = boto3.client("s3")

    keys = getpath(s3,provider,date)

    articles = []

    for key in keys:
        try: 
            file = s3.get_object(Bucket="nii-dev",Key=key)
            content = file["Body"].read().decode("utf-8")
            articleList = re.split(r'\.|\,|\s|\t',content)
        except Exception as e:
            print("[ERROR]FileOpenError: " + str(e))
        
        articles += [re.sub('\"|\!|\?',"",a.lower()) for a in articleList]
    
    df = pd.DataFrame(articles,columns=["word"])

    df_groupby = df.groupby(["word"]).size()

    output(df_groupby,provider,date)

    # for i in range(0,20):
    #     index = df_groupby.index[i]
    #     values = df_groupby.values[i]
    #     print(index,values)

    # local_output(df_groupby,provider,day)

    # print(df_groupby.count())
    

    # S3オブジェクト取得テスト
    # s3 = boto3.client("s3")

    # file = s3.get_object(Bucket="nii-dev",Key="english-frequency/articles/BSS/2023-09-03/5.txt")
    # print(file)

    # db接続テスト
    # with pymysql.connect(host=ENDPOINT, user=USER, passwd=os.environ["db_pass"], db=DBNAME) as conn:
    #     with conn.cursor() as cur:
    #         cur.execute("select * from mst_wordtype;")
    #         for r in cur:
    #             print(r)