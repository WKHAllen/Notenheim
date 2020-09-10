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

### Get profile info

* Path: `/api/getProfileInfo`
* Returns: `info: { email (string), joinTimestamp (int) }`

### Change password

* Path: `/api/changePassword`
* Parameters: `oldPassword` (string), `newPassword` (string)

### New list

* Path: `/api/newList`
* Parameters: `title` (string)
* Returns: `listID` (string)

### Get lists

* Path: `/api/getLists`
* Returns: `lists: [ { listID (string), title (string), updateTimestamp (int) }, ... ]`

### List info

* Path: `/api/listInfo`
* Parameters: `listID` (string)
* Returns: `info: { title (string), items: [ { listItemID (string), content (string), position (int), checked (boolean) }, ... ] }`

### Rename list

* Path: `/api/renameList`
* Parameters: `listID` (string), `newName` (string)

### Delete list

* Path: `/api/deleteList`
* Parameters: `listID` (string)

### New list item

* Path: `/api/newListItem`
* Parameters: `listID` (string)
* Returns: `listItemID` (string)

### Edit list item

* Path: `/api/editListItem`
* Parameters: `listItemID` (string), `newContent` (string)

### Check list item

* Path: `/api/checkListItem`
* Parameters: `listItemID` (string), `checked` (boolean)

### Move list item

* Path `/api/moveListItem`
* Parameters: `listItemID` (string), `newPosition` (int)

### Delete list item

* Path `/api/deleteListItem`
* Parameters: `listItemID` (string)
