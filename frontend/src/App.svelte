<script>
    import { DownloadVideos, SetOutputDir } from "../wailsjs/go/main/App.js";

    let directory = "";
    let url;

    function download() {
        DownloadVideos(url).then((result) => (console.log(result)));
    }

    function setOutputDir() {
        SetOutputDir().then((result) => (directory = result));
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
                    bind:value={url}
                    placeholder="https://erome.com/..."
                />
            </label>
            </div>
            <div class="grid">
            <div class="grid">
                <label for="Output Dir">
                    Output Dir
                    <input
                        autocomplete="off"
                        bind:value={directory}
                        placeholder="~/Downloads"
                        class="input"
                        id="name"
                        type="text"
                    />
                </label>
            </div>
        </div>
        <div class="grid">
        <button on:click={setOutputDir} class="secondary">Choose Output Dir</button>
        <button on:click={download}>Download</button>
        </div>
    </section>
</main>

<style>
    #main-section {
        width: 80%;
    }
</style>
