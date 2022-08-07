<script lang="ts">
	// import { spring } from 'svelte/motion';

	export let assignee = '';
	export let IsDone = false;
	export let Detail = '';
	export let deadline = '';
	export let onClickDetail = () => {};
	export let onClickDone = (status: boolean) => {};
	export let onClickDelete = () => {};
	function handleClickCheckbox() {
		IsDone = !IsDone;
		onClickDone(IsDone);
	}

	// convert rfc3339 to date Ymd
	function convertRfc3339ToDate(rfc3339: string) {
		const date = new Date(rfc3339);
		const Y = date.getFullYear();
		const M = date.getMonth() + 1;
		const D = date.getDate();
		return `${Y}-${M}-${D}`;
	}

	let dateString = convertRfc3339ToDate(deadline);
</script>

<div class="input-group mb-3">
	<div class="input-group-text">
		<input
			on:click={handleClickCheckbox}
			checked={IsDone}
			class="form-check-input mt-0"
			type="checkbox"
			value=""
			aria-label="Checkbox for following text input"
		/>
	</div>
	<div class="form-control" on:dblclick={() => onClickDetail()}>{Detail}</div>
	<small class="input-group-text text-muted">{dateString}</small>
	<small class="input-group-text text-muted">{assignee}</small>
	<button on:click={onClickDelete} class="input-group-text text-bg-danger pe-auto">X</button>
</div>
