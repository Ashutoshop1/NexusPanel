package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetMetrics returns monitoring metrics
func GetMetrics(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "Get metrics not yet implemented",
	})
}

// ListAlerts returns list of alerts
func ListAlerts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": []interface{}{},
		"message": "Alert list (not yet implemented)",
	})
}

// AcknowledgeAlert acknowledges an alert
func AcknowledgeAlert(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "Acknowledge alert not yet implemented",
	})
}

// ResolveAlert resolves an alert
func ResolveAlert(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "Resolve alert not yet implemented",
	})
}

// ListAlertRules returns list of alert rules
func ListAlertRules(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": []interface{}{},
		"message": "Alert rule list (not yet implemented)",
	})
}

// CreateAlertRule creates a new alert rule
func CreateAlertRule(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "Create alert rule not yet implemented",
	})
}

// GetAlertRule returns alert rule details
func GetAlertRule(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "Get alert rule not yet implemented",
	})
}

// UpdateAlertRule updates an alert rule
func UpdateAlertRule(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "Update alert rule not yet implemented",
	})
}

// DeleteAlertRule deletes an alert rule
func DeleteAlertRule(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "Delete alert rule not yet implemented",
	})
}

// WebSocketTerminal handles WebSocket terminal connections
func WebSocketTerminal(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "WebSocket terminal not yet implemented",
	})
}

// WebSocketMetrics handles WebSocket metrics connections
func WebSocketMetrics(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "WebSocket metrics not yet implemented",
	})
}
