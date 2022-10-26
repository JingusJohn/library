<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	// dispatcher for letting parent component know that a new component has been created
	const dispatch = createEventDispatcher();

	let title: string = '';
	let author: string = '';
	let description: string = '';

	function submit(): void {
		if (validStringAnswers([title, author])) {
			// do submit
			dispatch('bookCreation', {
				bookCreated: true,
				bookData: {
					title: title,
					author: author,
					description: description
				}
			});
		}
	}

	function validStringAnswers(answers: string[]): boolean {
		// might need to add check for already existing books
		for (let i of answers) {
			if (i == null || i == '') {
				return false;
			}
		}
		return true;
	}
</script>

<div class="form-control flex flex-col items-center space-y-4">

	<input placeholder="Title" class="input-bordered input w-full max-w-xs" id="title-input" bind:value={title} />

	<input placeholder="Author" class="input-bordered input w-full max-w-xs" bind:value={author} />

	<input placeholder="Description" class="input-bordered input w-full max-w-xs" bind:value={description} />

	<div class="modal-action">
		<button
			for="my-modal-4"
			class="btn"
			disabled={!validStringAnswers([title, author, description])}
			on:click={submit}>Create</button
		>
	</div>
</div>
