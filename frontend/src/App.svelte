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
        </div>
        <button on:click={download}>Download</button>
    </section>
</main>

<style>
    #main-section {
        width: 50%;
    }
</style>
