<script lang="ts">
	import NewBook from '../../components/library/home/NewBook.svelte';
	import type { Book } from '../../lib/models/book';

	export let data;

	let showNewBook: boolean = false;
	let { books } = data;

	// event handlers
	function handleNewBook(ev: string) {
		if (ev == 'create') {
			showNewBook = true;
		} else if (ev == 'cancel') {
			showNewBook = false;
		}
	}

	async function handleNewBookMessage(event: any) {
		if (event.detail.bookCreated) {
			showNewBook = false;
			const data = JSON.stringify(event.detail.bookData);
			const headers = {
				'Content-Type': 'application/json'
			};
			const res = await fetch('http://localhost:8080/create-book', {
				method: 'POST',
				headers,
				body: data
			})
				.then((response) => response.json())
				.then((data) => {
					console.log(data);
					let newBook: Book = data;
					books.push(newBook);
					books = books;
				});
		}
	}
</script>

<div class="flex flex-col">
	<div class="items-center">
		<h1 class="text-center text-xl">My Library</h1>

		<label for="my-modal" class="modal-button btn">Add Book</label>

		<!-- Put this part before </body> tag -->
		<input type="checkbox" id="my-modal" class="modal-toggle" bind:checked={showNewBook} />
		<div class="modal" on:click|self={() => (showNewBook = false)}>
			<div class="modal-box">
				<h3 class="text-lg font-bold text-center">Add a book to your library</h3>
				<NewBook on:bookCreation={handleNewBookMessage} />
			</div>
		</div>
	</div>

	<!-- The button to open modal -->

	<div class="overflow-x-auto">
		<table class="table w-full">
			<tr class="bg-base-200">
				<th>Title</th>
				<th>Author</th>
				<th>Description</th>
			</tr>

			{#each books as b}
				<tr class="hover">
					<td><a href="/library/book/{b.id}">{b.title}</a></td>
					<td>{b.author}</td>
					<td>{b.description}</td>
				</tr>
			{/each}
		</table>
	</div>
</div>
