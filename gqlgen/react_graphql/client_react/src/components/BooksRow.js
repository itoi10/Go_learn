
const BooksRow = (props) => {
  return (
    <tr>
      <td>{props.isbn}</td>
      <td>{props.title}</td>
      <td>{props.authors.join(', ')}</td>
      <td>{props.price}</td>
    </tr>
  )
}

export default BooksRow