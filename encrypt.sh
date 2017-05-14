mkdir /tmp/auth-go
git clone github.com/mrlsd/gpg-test /tmp/auth-go/
tar -cvpzf ../auth-kit /tmp/auth-go/auth-go.tar.gz
gpg -e -r mrlsd@ya.ru /tmp/auth-go/auth-go.tar.gz
rm -f /tmp/auth-go/auth-go.tar.gz
св /tmp/auth-go/
git add .
git commit -a
git push -u origin master