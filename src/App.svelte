<script>
	var germantogreek = false;
	var inputplaceholder = "Griechische Vokabel angeben";
	var displaykeyboard = false;
	var displaykeyboard_css = "none";
	function showdisplaykeyboard() {
		if (displaykeyboard == true) {
			displaykeyboard_css = "none";
			displaykeyboard = false;
		} else {
			displaykeyboard_css = "block";
			displaykeyboard = true;
		}
	}
	function change_germantogreek() {
		germantogreek = true;
		inputplaceholder = "Deutsche Vokabel angeben";
	}

	function change_greektogerman() {
		germantogreek = false;
		inputplaceholder = "Griechische Vokabel angeben";
	}
</script>

<main>
	<div class="search">
		<div class="trans_switcher">
			<button on:click={() => change_greektogerman()}
				>Griechisch zu Deutsch</button
			>
			<button on:click={() => change_germantogreek()}
				>Deutsch zu Griechisch</button
			>
		</div>
		<input type="text" placeholder={inputplaceholder} />
		<button>Suchen</button>
		<button on:click={() => showdisplaykeyboard()}
			>Griechische Tastatur anzeigen</button
		>
	</div>
	<div class="keyboard" style="display: {displaykeyboard_css};">
		<button>dwdw</button>
	</div>
</main>

<style>
	main {
		display: flex;
		flex-direction: column;
		height: 100vh;
		width: 100vw;
		justify-content: center;
		align-items: center;
	}

	.search {
		display: flex;
		flex-direction: column;
	}
	button,
	input {
		margin: 0;
	}
</style>
