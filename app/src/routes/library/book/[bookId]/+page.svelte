<script lang="ts">
  
  interface Book {
    id?: string,
    title: string,
    author: string,
    description: string
  }

  let book: Book;

  export let data;
  book = data.book;
</script>

<h1>{book.title}</h1>
<h2>{book.author}</h2>
<p>{book.description}</p>
