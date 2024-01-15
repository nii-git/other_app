import glob,sys,re
import pandas as pd

def local_init():
    if len(sys.argv) != 3:
        print("[ERROR] illegal commandline argument")
        sys.exit(-1)
        
    provider = sys.argv[1]
    day = sys.argv[2]
    return provider,day

def local_getpath(path):
    files = glob.glob(path)
    return files

def local_output(df_groupby,provider,day):
    with open("./result/" + provider + "/" + day + "/result.csv","a") as f:
        for i in range(0,df_groupby.count()):
            index = df_groupby.index[i]
            values = df_groupby.values[i]
            f.write(index + "," + str(values) + "\n")




if __name__ == "__main__":
    provider,day = local_init()
    # S3 -> Lambda 1日分の英語ファイルを取得する
    targetPaths = local_getpath("./output/" + provider + "/" + day + "/*")
    print(targetPaths)

    articles = []

    for t in targetPaths:
        try: 
            print(t)
            with open(t,"r") as f:
                #エラー出るならサイズを分ける
                articleList = re.split(r'\.|\,|\s|\t',f.read(-1))
        except Exception as e:
            print("[ERROR]FileOpenError: " + str(e))
        
        articles += [re.sub('\"|\!|\?',"",a.lower()) for a in articleList]
    
    df = pd.DataFrame(articles,columns=["word"])

    df_groupby = df.groupby(["word"]).size()

    local_output(df_groupby,provider,day)

    # print(df_groupby.count())
    