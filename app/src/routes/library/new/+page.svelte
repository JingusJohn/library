<script lang="ts">
  let title: string = "";
  let author: string = "";
  let message: string = "";


  function submit(): void {
    if (validStringAnswers([title, author])) {
      // do submit
      message = "Book Created!"
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

<label>
  Title
  <input id="title-input" bind:value={title}>
</label>

<label>
  Author
  <input bind:value={author}>
</label>

<button disabled='{!validStringAnswers([title, author])}' on:click={submit}>Create</button>

<p>{message}</p>
