rm -rf /tmp/auth-kit/
mkdir /tmp/auth-kit/
git clone github.com/mrlsd/gpg-test /tmp/auth-kit/
cp ./ /tmp/auth-kit/
tar -cvzf /tmp/auth-kit/auth-kit.tar.gz ../auth-kit
gpg -e -r mrlsd@ya.ru /tmp/auth-kit/auth-kit.tar.gz
rm -f /tmp/auth-kit/auth-kit.tar.gz
cd /tmp/auth-kit/
git add .
git commit -a
## git push -u origin master