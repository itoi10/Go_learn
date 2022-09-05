import BooksRow from './BooksRow'
import fetchGraphQL from '../api/fetchGraphQL'
import React, { useState, useEffect } from 'react'

const BooksTable = () => {

  const [books, setBooks] = useState([]);

  // GraphQLクエリ
  const query = `query getBooks($title: String!) {
    Books(title: $title) {
      isbn
      title
      authors
      price
    }
  }`;

  useEffect(() => {
    // GraphQLサーバから情報取得
    fetchGraphQL(query, { title: 'a' })
      .then(response => {
        setBooks(response.data.Books);
      }).catch(error => {
        console.error(error);
      });
  });

  const rows = [];
  books.forEach(book => {
    rows.push(
      <BooksRow
        isbn={book.isbn}
        title={book.title}
        authors={book.authors}
        price={book.price}
        key={book.isbn}
      />
    );
  });

  return (
    <table>
      <thead>
        <tr>
          <th>ISBN</th>
          <th>Title</th>
          <th>Authors</th>
          <th>Price</th>
        </tr>
      </thead>
      <tbody>
        {rows}
      </tbody>
    </table>
  )
}

export default BooksTable