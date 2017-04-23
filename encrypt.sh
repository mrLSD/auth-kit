rm -rf /tmp/auth-kit/
cp -r ./ /tmp/auth-kit-tmp/
git clone git@github.com:mrLSD/gpg-test.git /tmp/auth-kit/
tar -cvzf /tmp/auth-kit/auth-kit.tar.gz /tmp/auth-kit-tmp/
rm -rf /tmp/auth-kit-tmp/
gpg -e -r mrlsd@ya.ru /tmp/auth-kit/auth-kit.tar.gz
rm -f /tmp/auth-kit/auth-kit.tar.gz
cd /tmp/auth-kit/
git add .
git commit -a --date="1 days ago"
# git push -u origin master
