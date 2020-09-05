# Notenheim

Notenheim is an app that helps you keep track of all your notes and lists from one place. It can be accessed across all devices, and features an exotic design.

## Deployment

The app is deployed to [notenheim.com](https://www.notenheim.com/).

## Stack

For this project, the [PGTR](https://github.com/WKHAllen/pgtr-sample-app) stack is used.

## REST API Endpoints

Listed below are all the available REST endpoints. All endpoints contain the attribute `error` which will be an error string if an error occurred, or `null` if not.

### Register

* Path: `/api/register`
* Parameters: `email` (string), `password` (string)

### Login

* Path: `/api/login`
* Parameters: `email` (string), `password` (string)

### Logout

* Path: `/api/logout`

### Verification

* Path: `/api/verify`
* Parameters: `verifyID` (string)

### Request password reset

* Path: `/api/requestPasswordReset`
* Parameters: `email` (string)

### Check password reset ID validity

* Path: `/api/validPasswordResetID`
* Parameters: `resetID` (string)

### Reset password

* Path: `/api/resetPassword`
* Parameters: `resetID` (string), `newPassword` (string)

### Change password

* Path: `/api/changePassword`
* Parameters: `newPassword` (string)
