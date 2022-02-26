# DirectoryHashCalculator
ディレクトリのハッシュ計算

# 実行方法
DirectoryHash [--hashAlgorithm=${hashAlgorithm}] [targetDirectory]  

hashAlgorithmには次のハッシュ関数を指定する
- MD5: デフォルト

targetDirectoryには対象となるディレクトリを指定する（デフォルトはカレントディレクトリ）

# アルゴリズム
要素  
1バイト目
- 0x00: ディレクトリ
- 0x01: ファイル
2バイト目以降ファイル名 + ハッシュ値  

各要素を0x00で結合する
