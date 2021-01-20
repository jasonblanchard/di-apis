import { NotebookClient, RequestHeadersHandler } from "../client"

const notebookClient = new NotebookClient("https://di3.blanktech.net/notebook")

notebookClient.setRequestHeadersHandler((headers) => ({
    ...headers,
    "Authorization": `Bearer ${process.env.API_TOKEN}`,
}))

notebookClient.Notebook_GetEntry({
    id: "8351",
}).then(response => {
    console.log("Body: ", response.body)
}).catch(error => {
    console.log("Error: ", error.status, error.response.body)
})