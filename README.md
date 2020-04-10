# Alg_mkdirfile
매번 폴더랑 파일 만들기 귀찮아서 만든 프로그램

## How to use
```bash
.\main.exe -s Beakjoon -d Bronze4 -p ASDF -l go
```
만약 띄어쓰기가 있다면 ""로 감싸준다
```bash
.\main.exe -s Beakjoon -d Bronze4 -p "현실에선 초등학교 교사 이세계에선 Gopher" -l go
```


## Param
```bash
Usage: main.exe --site SITE --diff DIFF --prob PROB --lang LANG

Options:
  --site SITE, -s SITE   Site name
  --diff DIFF, -d DIFF   Difficulty
  --prob PROB, -p PROB   Problem Name
  --lang LANG, -l LANG   Language Name
  --help, -h             display this help and exit
```


## TODO
1. 난이도 입력 간소화
2. 언어별로 처리 단계 최적화 
3. default.x 파일로 기본틀 제공