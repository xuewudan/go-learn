# 用户注册
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{"username":"test","password":"123456","email":"test@example.com"}'

# 用户登录
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"username":"test","password":"123456"}'

# 创建文章（需替换 token）
curl -X POST http://localhost:8080/api/posts \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6InRlc3QiLCJpc3MiOiJibG9nLWJhY2tlbmQiLCJleHAiOjE3NjE2MzU4NjksIm5iZiI6MTc2MTU0OTQ2OSwiaWF0IjoxNzYxNTQ5NDY5fQ.qhp9X5966vOGnh3_l1Amui97EawARt5oDtCqbzWv3i8" \
  -d '{"title":"测试文章","content":"这是一篇测试文章"}'