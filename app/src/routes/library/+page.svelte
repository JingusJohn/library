<script lang="ts">
  import { onMount } from "svelte";
	import NewBook from "../../components/library/home/NewBook.svelte";

  interface Book {
    title: string,
    author: string,
    description: string
  }
  
  let showNewBook: boolean = false;
  let msg: string = "";

  onMount(async () => {
    fetch("http://localhost:8080/helloWorld")
      .then(response => response.json())
      .then(data => {
        console.log(data)
        msg = data.msg;
      }).catch(error => {
        console.log(error);
        return [];
      });
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
      msg = event.detail.bookData.Title + " created!"
      console.log(event.detail.bookData)
      console.log(JSON.stringify(event.detail.bookData))
      const res = await fetch("http://localhost:8080/create-book/", {
        method: 'POST',
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({
          title: event.detail.bookData.title,
          author: event.detail.bookData.author,
          description: event.detail.bookData.description
        })
      }).then(response => response.json())
        .then(data => {
          console.log(data)
        })

      // result = JSON.stringify(json)
    }


  }

</script>

<div id="body" class="body">

</div>

<h1>The Library</h1>
<p>Here, you can find every book in my personal library!</p>

{#if !showNewBook}
<button on:click={() => handleNewBook("create")}>New Book</button>
{:else}
<button on:click={() => handleNewBook("cancel")}>Cancel</button>
{/if}

{#if showNewBook}
<NewBook on:bookCreation={handleNewBookMessage}/>
{/if}

<p>{msg}</p>

<style>
  .body {
    display: flex;
    flex-direction: column;
  }
</style>
