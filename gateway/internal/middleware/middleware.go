package middleware

// AuthMiddleware is a middleware function for validating company tokens.
//func AuthMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		resp := types.Result{}
//
//		// Retrieve Token from the request header.
//		token := c.GetHeader("Token")
//		if token == "" {
//			// If the Token is empty, return an authentication failure error response.
//			c.JSON(http.StatusOK, resp.Fail(ecode.AuthError))
//			c.Abort()
//			return
//		}
//
//		// Create a request to check the company token.
//		req := &company_v1.CheckCompanyTokenRequest{
//			Token: token,
//		}
//
//		// Call the CheckCompanyToken method of the company service.
//		info, err := company.CheckCompanyToken(context.Background(), req)
//
//		// If an error occurs, return an authentication failure error response.
//		if err != nil {
//			c.JSON(http.StatusOK, resp.Fail(ecode.AuthError))
//			c.Abort()
//			return
//		}
//
//		// Store the login information in the Gin context for later use in subsequent processing.
//		c.Set("LoginInfo", info)
//	}
//}
