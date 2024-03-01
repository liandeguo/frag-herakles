<script>
	var germantogreek = false;
	var inputplaceholder = "Griechische Vokabel angeben";
	var displaykeyboard = false;
	var displaykeyboard_css = "none";
	var input = "";
	function search() {
		console.log(input);
		const url = "127.0.0.1/answer/" + input;
		const options = {
			method: "POST", // Replace with desired method (GET, POST, etc.)
			// Add other options like headers, body, etc. as needed
		};

		async function makeRequest() {
			try {
				const response = await fetch(url, options);
				if (!response.ok) {
					throw new Error(`Error: ${response.status}`);
				}
				const data = await response.json();
				console.log(data); // Log the response data
			} catch (error) {
				console.error("Error:", error);
			}
		}

		makeRequest();
	}
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
		<input type="text" placeholder={inputplaceholder} bind:value={input} />
		<button on:click={() => search()}>Suchen</button>
		<button on:click={() => showdisplaykeyboard()}
			>Griechische Tastatur anzeigen</button
		>
	</div>
	<div class="keyboard" style="display: {displaykeyboard_css};">
		<button>dwdw</button>
		<a href="/word">Hello World!</a>
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
