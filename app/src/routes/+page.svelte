<script lang="ts">
	export let data;
</script>

<div class="flex justify-center items-center h-screen">
	<h1>Welcome, {data?.profile?.name || ''}!</h1>
</div>
