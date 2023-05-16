# 设置执行策略为 RemoteSigned
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser

# 执行你的项目构建命令
git pull
cd actor
npm run build
cd ..
go install