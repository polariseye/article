rem hugo server --theme=hyde --buildDrafts

hugo.exe --buildDrafts --baseUrl="https://polariseye.coding.me/" -d ../polariseye.coding.me
hugo.exe --buildDrafts --baseUrl="https://polariseye.github.io/" -d ../polariseye.github.io

echo "开始更新polariseye.github.io"
cd ../polariseye.github.io
git add .
git commit -am '更新blog'
git push origin master

echo "开始更新polariseye.coding.me"
cd ../polariseye.coding.me
git add .
git commit -am '更新blog'
git push origin master

pause