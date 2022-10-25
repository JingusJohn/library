<script lang="ts">
	import { createEventDispatcher } from "svelte";

  // dispatcher for letting parent component know that a new component has been created
  const dispatch = createEventDispatcher();

  let title: string = "";
  let author: string = "";
  let description: string = ""
  let message: string = "";


  function submit(): void {
    if (validStringAnswers([title, author])) {
      // do submit
      message = "Book Created!";
      dispatch('bookCreation', {
        bookCreated: true,
        bookData: {
          title: title,
          author: author,
          description: description
        }
      })
    }
  }

  function validStringAnswers(answers: string[]): boolean {
    // might need to add check for already existing books
    for (let i of answers) {
      if (i == null || i == "") {
        return false;
      }
    }
    return true;
  }
</script>

<div class="items-center h-screen">

  <label>
    Title
    <input class="input input-bordered w-full max-w-xs" id="title-input" bind:value={title}>
  </label>

  <label>
    Author
    <input class="input input-bordered w-full max-w-xs" bind:value={author}>
  </label>

  <label>
    Description
    <input class="input input-bordered w-full max-w-xs" bind:value={description}>
  </label>


  <button class="btn" disabled='{!validStringAnswers([title, author, description])}' on:click={submit}>Create</button>

  <p>{message}</p>

</div>

<style>
  .body {
    padding-top: 10px;
    display: flex;
    flex-direction: row;
  }
</style>
