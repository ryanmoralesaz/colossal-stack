# About

This is a practice development project utilizing Go, Fiber, GORM, PostgreSQL, and GraphQL.

## Author
Ryan Morales

# GraphQL Playground Commands Playbook

Open: **https://api.ryanmorales.info/colossal/playground**

---

## 1. Query All Books (Public - No Auth)

```graphql
query GetAllBooks {
  books {
    id
    title
    author
    publisher
    createdAt
    updatedAt
  }
}
```

---

## 2. Query Single Book by ID (Public - No Auth)

```graphql
query GetBook {
  book(id: "1") {
    id
    title
    author
    publisher
    createdAt
    updatedAt
  }
}
```

---

## 3. Create Book (Requires Auth)

**First, add HTTP header (click "HTTP HEADERS" at bottom):**

```json
{
  "Authorization": "Bearer YOUR_TOKEN_FROM_LOGIN"
}
```

**Then run:**

```graphql
mutation CreateBook {
  createBook(input: {
    title: "Learning GraphQL"
    author: "Ryan Morales"
    publisher: "Tech Press"
  }) {
    id
    title
    author
    publisher
    createdAt
  }
}
```

---

## 4. Update Book (Requires Auth)

```graphql
mutation UpdateBook {
  updateBook(id: "1", input: {
    title: "Mastering GraphQL"
    publisher: "Advanced Tech Press"
  }) {
    id
    title
    author
    publisher
    updatedAt
  }
}
```

**Note:** Only updates the fields you provide. Author stays the same in this example.

---

## 5. Delete Book (Requires Auth)

```graphql
mutation DeleteBook {
  deleteBook(id: "1")
}
```

**Returns:** `true` if successful

---

## 6. Create Multiple Books (Variables)

**In "Query Variables" section (bottom left):**

```json
{
  "book1": {
    "title": "Go Programming",
    "author": "Ryan Morales",
    "publisher": "Dev Books"
  },
  "book2": {
    "title": "PostgreSQL Mastery",
    "author": "Ryan Morales"
  }
}
```

**Query:**

```graphql
mutation CreateMultiple($book1: CreateBookInput!, $book2: CreateBookInput!) {
  first: createBook(input: $book1) {
    id
    title
  }
  second: createBook(input: $book2) {
    id
    title
  }
}
```

---

## 7. Test Auth Failure (No Token)

**Remove the Authorization header, then:**

```graphql
mutation ShouldFail {
  createBook(input: {
    title: "Unauthorized"
    author: "Hacker"
  }) {
    id
  }
}
```

**Expected Error:**
```json
{
  "errors": [
    {
      "message": "unauthorized: authentication required for this operation"
    }
  ]
}
```

---

## 8. Complete CRUD Demo

```graphql
# Step 1: Create
mutation {
  createBook(input: {
    title: "Demo Book"
    author: "Ryan Morales"
    publisher: "Demo Press"
  }) {
    id
    title
  }
}

# Step 2: Read (use ID from step 1)
query {
  book(id: "INSERT_ID_HERE") {
    id
    title
    author
    publisher
  }
}

# Step 3: Update
mutation {
  updateBook(id: "INSERT_ID_HERE", input: {
    title: "Updated Demo Book"
  }) {
    id
    title
    updatedAt
  }
}

# Step 4: Read again (verify update)
query {
  book(id: "INSERT_ID_HERE") {
    id
    title
    updatedAt
  }
}

# Step 5: Delete
mutation {
  deleteBook(id: "INSERT_ID_HERE")
}

# Step 6: Verify deletion
query {
  book(id: "INSERT_ID_HERE") {
    id
  }
}
```

---

## Getting Your Auth Token

**Method 1: Using kulala in Neovim**
- Open `api_test.http`
- Run the "Login" request with `<leader>rr`
- Copy token from response

**Method 2: Using curl**

```bash
curl -X POST https://api.ryanmorales.info/colossal/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "ryan@test.com", "password": "SecurePass123"}' | jq -r '.token'
```

Copy the output and paste into HTTP HEADERS as:

```json
{
  "Authorization": "Bearer PASTE_TOKEN_HERE"
}
```

---

## Pro Tips

1. **Use fragments for reusable fields:**

```graphql
fragment BookDetails on Book {
  id
  title
  author
  publisher
  createdAt
  updatedAt
}

query {
  books {
    ...BookDetails
  }
}
```

2. **Use aliases for multiple queries:**

```graphql
query {
  firstBook: book(id: "1") {
    title
  }
  secondBook: book(id: "2") {
    title
  }
}
```

3. **IntelliSense:** Press `Ctrl+Space` in the playground for autocomplete

4. **Docs:** Click "Schema" button on right side to see all available queries/mutations

---

**Save this for your portfolio README!**

