<script lang="ts">
	export let open = false;
	export let showBackdrop = true;
	export let onClosed = () => {};
	export let onOK = () => {};
	export let title = '';

	const modalClose = () => {
		open = false;
		if (onClosed) {
			onClosed();
		}
	};
</script>

{#if open}
	<div class="modal" tabindex="-1">
		<div class="modal-dialog">
			<div class="modal-content">
				<div class="modal-header">
					<h5 class="modal-title">{title}</h5>
					<button
						on:click={modalClose}
						type="button"
						class="btn-close"
						data-bs-dismiss="modal"
						aria-label="Close"
					/>
				</div>
				<div class="modal-body">
					<slot />
				</div>
				<div class="modal-footer">
					<button
						on:click={modalClose}
						type="button"
						class="btn btn-secondary"
						data-bs-dismiss="modal">Close</button
					>
					<button on:click={onOK} type="button" class="btn btn-primary">OK</button>
				</div>
			</div>
		</div>
	</div>
	{#if showBackdrop}
		<div class="modal-backdrop show" />
	{/if}
{/if}

<style>
	.modal {
		display: block;
	}
</style>
