rem hugo server --theme=hyde --buildDrafts

hugo.exe --buildDrafts --baseUrl="https://polariseye.coding.me/" -d ../polariseye.coding.me
hugo.exe --buildDrafts --baseUrl="https://polariseye.github.io/" -d ../polariseye.github.io

echo "��ʼ����polariseye.github.io"
cd ../polariseye.github.io
git add .
git commit -am '����blog'
git push origin master

echo "��ʼ����polariseye.coding.me"
cd ../polariseye.coding.me
git add .
git commit -am '����blog'
git push origin master

pause