import glob,sys,re,time

def local_getpath(path):
    files = glob.glob(path)
    return files



if __name__ == "__main__":

    # S3 -> Lambda 1日分の英語ファイルを取得する
    targetPaths = local_getpath("output/BSS/*/*")


    # targetPaths = local_getpath("output/BSS/2023-08-01/*")
    print(targetPaths)

    # 英字分割パターン
    pattern = re.compile(r'.|,|\s|\t')

    wordList = []

    for t in targetPaths:
        try: 
            print(t)
            with open(t,"r") as f:
                #エラー出るならサイズを分ける
                articleList = re.split(r'[.,\s]',f.read(-1))
                wordList += articleList
        except Exception as e:
            print("[ERROR]FileOpenError: " + str(e))
        
    print(len(wordList))



    # Test1:ニュース記事で比較

    # 全体に対して小文字化
    all_lower_start = time.time()
    all_lower_list = [w.lower() for w in wordList]
    all_lower_end = time.time()

    # print(all_lower_list)

    # チェックして小文字化
    check_islower_start = time.time()
    check_islower_list = [w if w.islower() else w.lower() for w in wordList]
    check_islower_end = time.time()

    # print(check_islower_list)
    print("Test1:ニュース記事で比較")
    print("all_lower:" + str(all_lower_end-all_lower_start))
    print("check_lower:" + str(check_islower_end-check_islower_start))

    print(100*((check_islower_end-check_islower_start)-(all_lower_end-all_lower_start))/(check_islower_end-check_islower_start)-(all_lower_end-all_lower_start))

    # Test2: 全て大文字だった場合の比較
    wordList = [w.upper() for w in wordList]


    # 全体に対して小文字化
    all_lower_start = time.time()
    all_lower_list = [w.lower() for w in wordList]
    all_lower_end = time.time()

    # チェックして小文字化
    check_islower_start = time.time()
    check_islower_list = [w if w.islower() else w.lower() for w in wordList]
    check_islower_end = time.time()

    print("Test2: 全て大文字だった場合の比較")
    print("all_lower:" + str(all_lower_end-all_lower_start))
    print("check_lower:" + str(check_islower_end-check_islower_start))


    print(100*((check_islower_end-check_islower_start)-(all_lower_end-all_lower_start))/(check_islower_end-check_islower_start)-(all_lower_end-all_lower_start))

    # Test3: 全て小文字だった場合の比較
    wordList = [w.lower() for w in wordList]


    # 全体に対して小文字化
    all_lower_start = time.time()
    all_lower_list = [w.lower() for w in wordList]
    all_lower_end = time.time()

    # チェックして小文字化
    check_islower_start = time.time()
    check_islower_list = [w if w.islower() else w.lower() for w in wordList]
    check_islower_end = time.time()

    print("Test3: 全て小文字だった場合の比較")
    print("all_lower:" + str(all_lower_end-all_lower_start))
    print("check_lower:" + str(check_islower_end-check_islower_start))


    print(100*((check_islower_end-check_islower_start)-(all_lower_end-all_lower_start))/(check_islower_end-check_islower_start)-(all_lower_end-all_lower_start))