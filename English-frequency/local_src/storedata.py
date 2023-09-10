import glob,sys,re

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



if __name__ == "__main__":
    provider,day = local_init()
    # S3 -> Lambda 1日分の英語ファイルを取得する
    targetPaths = local_getpath("./output/" + provider + "/" + day + "/*")
    print(targetPaths)

    # 英字分割パターン
    pattern = re.compile(r'.|,|\s|\t')

    for t in targetPaths:
        try: 
            print(t)
            with open(t,"r") as f:
                #エラー出るならサイズを分ける
                articleList = re.split(r'[.,\s]',f.read(-1))
        except Exception as e:
            print("[ERROR]FileOpenError: " + str(e))
        
        break