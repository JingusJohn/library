<script lang="ts">
  import { onMount } from "svelte";
	import NewBook from "../../components/library/home/NewBook.svelte";
  import type { Book } from "../../lib/models/book";


  export let data;
  
  let showNewBook: boolean = false;
  let books: Book[] = [];

  onMount(async () => {
    fetch("http://localhost:8080/get-books")
      .then(response => response.json())
      .then(data => {
        books = data;
      })
  });


  // event handlers
  function handleNewBook(ev: string) {
    if (ev == "create") {
      showNewBook = true;
    } else if (ev == "cancel") {
      showNewBook = false;
    }
  }

  async function handleNewBookMessage(event: any) {
    if (event.detail.bookCreated) {
      showNewBook = false;
      console.log(event.detail.bookData)
      console.log(JSON.stringify(event.detail.bookData))
      const data = JSON.stringify(event.detail.bookData);
      const headers = {
        'Content-Type': "application/json"
      }
      const res = await fetch("http://localhost:8080/create-book", {
        method: "POST",
        headers,
        body: data
      }).then(
        response => response.json()
      ).then(data => {
        console.log(data);
        let newBook: Book = data;
        books.push(newBook);
      })
    }

  }

</script>

<div id="body" class="body">

</div>

<h1 class="text-xl">My Library</h1>

{#if !showNewBook}
<button class="btn btn-primary btn-xs" on:click={() => handleNewBook("create")}>New Book</button>
{:else}
<button class="btn btn-warning btn-xs" on:click={() => handleNewBook("cancel")}>Cancel</button>
{/if}

{#if showNewBook}
<NewBook on:bookCreation={handleNewBookMessage}/>
{/if}

<table>
  <tr>
    <th>Title</th>
    <th>Author</th>
    <th>Description</th>
  </tr>

{#each books as b, i}
  <tr>
    <td><a href="/library/book/{b.id}">{b.title}</a></td>
    <td>{b.author}</td>
    <td>{b.description}</td>
  </tr>
{/each}

</table>


<style>
  .body {
    display: flex;
    flex-direction: column;
  }
</style>
