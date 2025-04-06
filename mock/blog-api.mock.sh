# 针对 blog 的几个 api，生成 curl 的相关 mock 请求 shell

# 获取所有文章
curl -X GET http://localhost:8889/api/v1/posts

# 获取单篇文章
curl -X GET http://localhost:8889/api/v1/posts/1

# 创建文章
curl -X POST http://localhost:8889/api/v1/posts \
  -H "Content-Type: application/json" \
  -d '{"title": "新文章", "content": "这是新文章的内容"}'

# 更新文章
curl -X PUT http://localhost:8889/api/v1/posts/新文章 \
  -H "Content-Type: application/json" \
  -d '{"title": "更新后的文章", "content": "这是更新后的文章内容"}'

# 删除文章
curl -X DELETE http://localhost:8889/api/v1/posts/新文章

# 创建一篇markdown格式的文章
curl -X POST http://localhost:8889/api/v1/posts \
  -H "Content-Type: application/json" \
  -d '{"title": "Markdown 示例文章", "content": "# 一级标题\n\n## 二级标题\n\n这是一段普通的文本，支持 **加粗**、*斜体* 等格式。\n\n### 列表示例\n\n- 项目1\n- 项目2\n- 项目3\n\n### 代码示例\n\n```python\ndef hello():\n    print(\"Hello World!\")\n```\n\n> 这是一段引用文本\n\n[这是一个链接](https://example.com)\n\n![图片描述](https://example.com/image.jpg)"}'


