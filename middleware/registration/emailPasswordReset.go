package registration

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"time"

	"github.com/sea350/ustart_go/middleware/client"

	get "github.com/sea350/ustart_go/get/user"
	post "github.com/sea350/ustart_go/post/user"
	uses "github.com/sea350/ustart_go/uses"
)

//SendPasswordResetEmail ... Sends password reset token link to user and saves token to their AuthenticationCode
//Requires a valid user email address
//Returns if there is an error
func SendPasswordResetEmail(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")

	var cs client.ClientSide

	defer client.RenderSidebar(w, r, "templateNoUser2")
	defer client.RenderTemplate(w, r, "reset-forgot-pw", cs)

	//If the email isn't blank and it is in use...
	if email != "" {
		emailInUse, err := get.EmailInUse(client.Eclient, email)
		if err != nil {
			fmt.Println("Error: ustart_go/middleware/registration/emailPasswordReset Line 30: Unable to retrieve email")
			fmt.Println(err)
		}

		if emailInUse {
			token, err := uses.GenerateRandomString(32)
			if err != nil {
				fmt.Println("Error ustart_go/middleware/registration/emailPasswordReset line 37: Error generating token")
				fmt.Println(err)
				return
			}

			userID, err := get.UserIDByEmail(client.Eclient, email)
			if err != nil {
				fmt.Println("Error ustart_go/middleware/registration/emailPasswordReset line 44: Unable to retreive userID by email")
				fmt.Println(err)
				return
			}

			err = post.UpdateUser(client.Eclient, userID, "AuthenticationCodeTime", time.Now())
			if err != nil {
				fmt.Println("Error ustart_go/middleware/registration/emailPasswordReset line 51: Error posting user")
				fmt.Println(err)
				return
			}

			err = post.UpdateUser(client.Eclient, userID, "AuthenticationCode", token)
			if err != nil {
				fmt.Println("Error ustart_go/middleware/registration/emailPasswordReset line 58: Error posting user")
				fmt.Println(err)
				return
			}

			//Todo: make from and pw not plaintext
			from := "ustarttestemail@gmail.com"
			pass := "Ust@rt20!8~~"
			body := `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
				<html xmlns="http://www.w3.org/1999/xhtml">
				<head>
				<title>USTART</title>
				<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
				<meta http-equiv="X-UA-Compatible" content="IE=edge" />
				<meta name="viewport" content="width=device-width, initial-scale=1.0 " />
				<style type="text/css">
					body {
					margin: 0 !important;
					padding: 0 !important;
					-webkit-text-size-adjust: 100% !important;
					-ms-text-size-adjust: 100% !important;
					-webkit-font-smoothing: antialiased !important;
					}
					img {
					border: 0 !important;
					outline: none !important;
					}
					p {
					Margin: 0px !important;
					Padding: 0px !important;
					}
					table {
					border-collapse: collapse;
					mso-table-lspace: 0px;
					mso-table-rspace: 0px;
					}
					td, a, span {
					border-collapse: collapse;
					mso-line-height-rule: exactly;
					}
					.ExternalClass * {
					line-height: 100%;
					}
					.em_defaultlink a {
					color: inherit !important;
					text-decoration: none !important;
					}
					span.MsoHyperlink {
					mso-style-priority: 99;
					color: inherit;
					}
					span.MsoHyperlinkFollowed {
					mso-style-priority: 99;
					color: inherit;
					}
					@media only screen and (min-width:481px) and (max-width:699px) {
					.em_main_table {
						width: 100% !important;
					}
					.em_wrapper {
						width: 100% !important;
					}
					.em_hide {
						display: none !important;
					}
					.em_img {
						width: 100% !important;
						height: auto !important;
					}
					.em_h20 {
						height: 20px !important;
					}
					.em_padd {
						padding: 20px 10px !important;
					}
					}
					@media screen and (max-width: 480px) {
					.em_main_table {
						width: 100% !important;
					}
					.em_wrapper {
						width: 100% !important;
					}
					.em_hide {
						display: none !important;
					}
					.em_img {
						width: 100% !important;
						height: auto !important;
					}
					.em_h20 {
						height: 20px !important;
					}
					.em_padd {
						padding: 20px 10px !important;
					}
					.em_text1 {
						font-size: 16px !important;
						line-height: 24px !important;
					}
					u + .em_body .em_full_wrap {
						width: 100% !important;
						width: 100vw !important;
					}
				}
				</style>
				</head>
				<body class="em_body" style="margin:0px; padding:0px;">
					<table class="em_full_wrap" valign="top" width="100%" cellspacing="0" cellpadding="0" border="0" bgcolor="#efefef" align="center">
					<tbody><tr>
						<td valign="top" align="center"><table class="em_main_table" style="width:700px;" width="700" cellspacing="0" cellpadding="0" border="0" align="center">
							<!--Top Banner-->
							<tr>
							<td valign="top" align="center"><table width="100%" cellspacing="0" cellpadding="0" border="0" align="center">
								<tbody><tr>
									<!-- fill src with banner img -->
									<td valign="top" align="center"><img class="em_img" style="display:block; font-family:Arial, sans-serif; font-size:30px; line-height:34px; max-width:700px;" src="" width="700" border="0" height="150"></td>
								</tr>
								</tbody></table></td>
							</tr>
							<!-- middle content -->
							<tr>
								<td valign="top" align="center" bgcolor="#fff" style="padding:35px 70px 30px;" class="em_padd"><table align="center" width="100%" border="0" cellspacing="0" cellpadding="0">
								<!-- title -->
								<tr>
									<td align="center" valign="top" style="font-family:'Open Sans', Arial, sans-serif; font-size:30px; line-height:30px;">Hi Username</td>
								</tr>
								<!-- white space-->
								<tr>
									<td height="15" style="font-size:0px; line-height:0px; height:15px;">&nbsp;</td>
								</tr>
								<!-- content -->
								<tr>
									<td valign="top" style="font-family:'Open Sans', Arial, sans-serif; font-size:16px; line-height:30px;">We received a request to reset your password for your Ustart Account. We would love to assist you!</td>
								</tr>
								<!-- white space -->
								<tr>
									<td height="15" style="font-size:0px; line-height:0px; height:20px;">&nbsp;</td>
								</tr>
								<!-- content -->
								<tr>
									<td valign="top" style="font-family:'Open Sans', Arial, sans-serif; font-size:16px; line-height:20px;">Simply click the button below to set a new password</td>
								</tr>
								<!-- white space -->
								<tr>
									<td height="15" style="font-size:0px; line-height:0px; height:50px;">&nbsp;</td>
								</tr>
								<!-- button -->
								<tr>
									<td valign="top" align ="center">
										<table cellspacing="0" cellpadding="0" border="0" align="center" style="width:25%;">
										<tbody>
											<tr>
											<!-- set button destination href here -->
											<td style="background-color: #4ecdc4; border-color: #4c5764; border-radius: 5px; padding: 10px;text-align: center;">
											<a style="display: block;color: #ffffff;font-size: 12px;text-decoration: none;text-transform: uppercase;" href="http://ustart.today:5002/ResetPassword/?email=` + email + `&verifCode=` + token + `">
												Change Pasword
												</a>
											</td>
											</tr>
										</tbody>
										</table>
									</td>
								</tr>
								<!-- white space -->
								<tr>
									<td height="15" style="font-size:0px; line-height:0px; height:40px;">&nbsp;</td>
								</tr>
								<!-- content -->
								<tr>
									<td valign="top" style="font-family:'Open Sans', Arial, sans-serif; font-size:16px; line-height:20px;">If you did not request a password reset, no further action is required.</td>
								</tr>
								</table></td>
							</tr>
							<!--Footer Banner-->
							<tr>
							<td style="padding:38px 30px;" class="em_padd" valign="top" bgcolor="#f6f7f8" align="center"><table width="100%" cellspacing="0" cellpadding="0" border="0" align="center">
								<tbody><tr>
									<td style="padding-bottom:16px;" valign="top" align="center"><table cellspacing="0" cellpadding="0" border="0" align="center">
										<tbody><tr>
										<!-- fill src with logo img -->
										<td valign="top" align="center"><a href="#" target="_blank" style="text-decoration:none;"><img src="" alt="fb" style="display:block; font-family:Arial, sans-serif; font-size:14px; line-height:14px; color:#ffffff; max-width:26px;" width="26" border="0" height="26"></a></td>
										<td style="width:6px;" width="6">&nbsp;</td>
										<td valign="top" align="center"><a href="#" target="_blank" style="text-decoration:none;"><img src="" alt="tw" style="display:block; font-family:Arial, sans-serif; font-size:14px; line-height:14px; color:#ffffff; max-width:27px;" width="27" border="0" height="26"></a></td>
										<td style="width:6px;" width="6">&nbsp;</td>
										<td valign="top" align="center"><a href="#" target="_blank" style="text-decoration:none;"><img src="" alt="insta" style="display:block; font-family:Arial, sans-serif; font-size:14px; line-height:14px; color:#ffffff; max-width:26px;" width="26" border="0" height="26"></a></td>
										</tr>
									</tbody></table></td>
								</tr>
								<tr>
									<td style="font-family:'Open Sans', Arial, sans-serif; font-size:11px; line-height:18px; color:#999999;" valign="top" align="center"><a href="#" target="_blank" style="color:#999999; text-decoration:underline;">PRIVACY STATEMENT</a> | <a href="#" target="_blank" style="color:#999999; text-decoration:underline;">TERMS OF SERVICE</a><br>
									© 2018 USTART. All Rights Not Reserved.<br>
									If you do not wish to receive any further emails from us,<a href="#" target="_blank" style="text-decoration:none; color:#999999;">Unsubscribe</a></td>
								</tr>
								</tbody></table></td>
							</tr>
						</table></td>
				</tr>
				</tbody></table>
					<div class="em_hide" style="white-space: nowrap; display: none; font-size:0px; line-height:0px;">&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;</div>
				</body>
				</html>
				
			`
			msg := "From: " + from + "\n" + "To: " + email + "\n" + "Subject: UStart Password Reset\n\n" + body

			err = smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, pass, "smtp.gmail.com"), from, []string{email}, []byte(msg))
			if err != nil {
				log.Printf("smtp error: %s", err)
				return
			}

		}
	}
}
