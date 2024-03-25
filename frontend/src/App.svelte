<script>
    import { DownloadVideos } from "../wailsjs/go/main/App.js";

    let resultText = "Enter the output dir";
    let name;

    function download() {
        DownloadVideos(name).then((result) => (resultText = result));
    }
</script>

<main>
    <section class="container" id="main-section">
        <div class="grid container">
            <label for="Url">
                Erome Url
                <input
                    type="text"
                    id="url"
                    name="Erome url"
                    placeholder="https://erome.com/..."
                />
            </label>
            <div class="grid">
                <label for="Output Dir">
                    Output Dir
                    <input
                        autocomplete="off"
                        bind:value={name}
                        placeholder="~/Downloads"
                        class="input"
                        id="name"
                        type="text"
                    />
                </label>
                <button on:click={download} class="secondary">Choose</button>

            </div>
        </div>
        <button on:click={download}>Download</button>
    </section>
</main>

<style>
    #main-section {
        width: 80%;
    }
</style>
