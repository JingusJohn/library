<script lang="ts">
	import NewBook from "../../components/library/home/NewBook.svelte";

  
  let showNewBook: boolean = false;
  let msg: string = "";


  // event handlers
  function handleNewBook(ev: string) {
    if (ev == "create") {
      showNewBook = true;
    } else if (ev == "cancel") {
      showNewBook = false;
    }
  }

  function handleNewBookMessage(event: any) {
    if (event.detail.bookCreated) {
      showNewBook = false;
      msg = event.detail.bookData.title + " created!"
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
