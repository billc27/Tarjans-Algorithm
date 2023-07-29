package routes

import (
    "os"
	"fmt"
	"tarjan-backend/utils"
    "github.com/gin-gonic/gin"
    "net/http"
)

func rootHandler(c *gin.Context) {
    c.String(http.StatusOK, "Hello!")
}

func RegisterRoutes(router *gin.Engine) {
    router.GET("/", rootHandler)
    router.POST("/visualize", visualizeHandler)
    router.POST("/scc", sccHandler)
    router.POST("/bridge", bridgeHandler)
}

func visualizeHandler(c *gin.Context) {
    // Get the data from the request body
    var requestData struct {
        Content string `json:"content"`
    }
	
    c.BindJSON(&requestData)

    // Access the data
    content := requestData.Content
    // fmt.Println(content)
    // fmt.Printf("%T\n", content)

	parsedContent := utils.ParseGraph(content)
    // fmt.Println(parsedContent)
    utils.DrawInputGraph(parsedContent)

    // Check if the file exists at the specified path
    filePath := "input-graph.gv.png"
    _, err := os.Stat(filePath)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "File not found"})
        return
    }

    // Return generated image in response
    c.File(filePath)
}

func sccHandler(c *gin.Context) {
    // Get the data from the request body
    var requestData struct {
        Content string `json:"content"`
    }
	fmt.Println(requestData.Content)
    c.BindJSON(&requestData)

    // Access the data
    content := requestData.Content

	parsedSCC := utils.ParseSCC(utils.GenerateSCC(content));
    utils.DrawSCCGraph(parsedSCC)

    filePath := "scc-graph.gv.png"
    _, err := os.Stat(filePath)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "File not found"})
        return
    }

    // Return generated image in response
    c.File(filePath)
}

func bridgeHandler(c *gin.Context) {
    // Get the data from the request body
    var requestData struct {
        Content string `json:"content"`
    }
	fmt.Println(requestData.Content)
    c.BindJSON(&requestData)

    // Access the data
    content := requestData.Content

	parsedBridge := utils.ParseGraph(utils.GenerateBridges(content))
    utils.DrawBridgeGraph(parsedBridge)

    filePath := "bridges-graph.png"
    _, err := os.Stat(filePath)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "File not found"})
        return
    }

    // Return generated image in response
    c.File(filePath)
}

