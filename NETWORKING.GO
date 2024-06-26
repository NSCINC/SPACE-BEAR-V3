{\rtf1\ansi\deff0\nouicompat{\fonttbl{\f0\fnil\fcharset0 Calibri;}}
{\*\generator Riched20 10.0.19041}\viewkind4\uc1 
\pard\sa200\sl276\slmult1\f0\fs22\lang22 Entendi. Vou criar um exemplo b\'e1sico de como voc\'ea pode criar um arquivo `index.go` em Go para interagir com servi\'e7os da AWS e Azure. Neste exemplo, usarei as bibliotecas SDK oficialmente fornecidas pela AWS e Azure para comunica\'e7\'e3o.\par
\par
```go\par
package main\par
\par
import (\par
    "fmt"\par
    "log"\par
    "net/http"\par
\par
    "github.com/aws/aws-sdk-go/aws"\par
    "github.com/aws/aws-sdk-go/aws/session"\par
    "github.com/aws/aws-sdk-go/service/s3"\par
\par
    "github.com/Azure/azure-sdk-for-go/storage"\par
)\par
\par
func main() \{\par
    // Configura o roteamento para diferentes endpoints\par
    http.HandleFunc("/aws", handleAWS)\par
    http.HandleFunc("/azure", handleAzure)\par
\par
    // Inicia o servidor na porta 8080\par
    fmt.Println("Server is running on localhost:8080")\par
    log.Fatal(http.ListenAndServe(":8080", nil))\par
\}\par
\par
func handleAWS(w http.ResponseWriter, r *http.Request) \{\par
    // Crie uma sess\'e3o AWS\par
    sess, err := session.NewSession(&aws.Config\{\par
        Region: aws.String("us-west-2"), // Altere a regi\'e3o conforme necess\'e1rio\par
    \})\par
    if err != nil \{\par
        http.Error(w, "Failed to create AWS session", http.StatusInternalServerError)\par
        return\par
    \}\par
\par
    // Crie um cliente S3\par
    svc := s3.New(sess)\par
\par
    // Exemplo de opera\'e7\'e3o S3\par
    _, err = svc.ListBuckets(nil)\par
    if err != nil \{\par
        http.Error(w, "Failed to list S3 buckets", http.StatusInternalServerError)\par
        return\par
    \}\par
\par
    // Opera\'e7\'e3o bem-sucedida\par
    fmt.Fprintf(w, "Interacted with AWS successfully")\par
\}\par
\par
func handleAzure(w http.ResponseWriter, r *http.Request) \{\par
    // Crie uma conta de armazenamento Azure\par
    accountName := "your-account-name"  // Insira o nome da sua conta de armazenamento Azure\par
    accountKey := "your-account-key"    // Insira a chave da sua conta de armazenamento Azure\par
\par
    // Crie um cliente de servi\'e7o de armazenamento Azure\par
    client, err := storage.NewBasicClient(accountName, accountKey)\par
    if err != nil \{\par
        http.Error(w, "Failed to create Azure storage client", http.StatusInternalServerError)\par
        return\par
    \}\par
\par
    // Obtenha o servi\'e7o de armazenamento\par
    service := client.GetBlobService()\par
\par
    // Exemplo de opera\'e7\'e3o de lista de cont\'eaineres do Azure Blob Storage\par
    containers, err := service.ListContainers(storage.ListContainersParameters\{\})\par
    if err != nil \{\par
        http.Error(w, "Failed to list Azure Blob Storage containers", http.StatusInternalServerError)\par
        return\par
    \}\par
\par
    // Opera\'e7\'e3o bem-sucedida\par
    fmt.Fprintf(w, "Interacted with Azure successfully")\par
\}\par
```\par
\par
Certifique-se de substituir `"your-account-name"` e `"your-account-key"` pelos valores reais da sua conta de armazenamento Azure. Al\'e9m disso, ajuste a regi\'e3o AWS conforme necess\'e1rio na fun\'e7\'e3o `handleAWS`.\par
}
 
