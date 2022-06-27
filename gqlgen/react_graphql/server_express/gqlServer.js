const app = require('express')();
const { graphqlHTTP } = require('express-graphql');
const { buildSchema } = require('graphql')
const cors = require('cors');


// データ
const books = [
  {
    isbn: '123-001',
    title: 'GraphQL Sample Book 1',
    authors: ['Remu', 'Ramu'],
    price: 4.98,
  },
  {
    isbn: '123-002',
    title: 'GraphQL Sample Book 2',
    authors: ['Taro', 'Jiro'],
    price: 1.02,
  },
  {
    isbn: '123-003',
    title: 'GraphQL Sample Book 3',
    authors: ['Alice'],
    price: 17.29,
  },
]

// スキーマ
const bookSchema = buildSchema(`
  type Query {
    Book(isbn: String!): Book
    Books(title: String): [Book]
  },
  type Book {
    isbn: String
    title: String
    authors: [String]
    price: Float
  }
`)

// クエリ
const BookQuery = args => {
  const isbn = args.isbn;
  return books.filter(Book => Book.isbn === isbn)[0]
}

const BooksQuery = args => {
  if (args.title) {
    const title = args.title;
    return books.filter(Book => Book.title.match(new RegExp(title)));
  }
  return books;
}


app.use(cors())
app.use('/graphql', graphqlHTTP({
  schema: bookSchema,
  rootValue: {
    Book: BookQuery,
    Books: BooksQuery,
  },
  graphiql: true,
}))


const defaultPort = 8080

app.listen(defaultPort, () => {
  console.log(`GraphQL Server is running on localhost://${defaultPort}/graphql`)
})