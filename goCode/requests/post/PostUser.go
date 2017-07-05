package post

import(
    elastic "gopkg.in/olivere/elastic.v5"
    types "github.com/nicolaifsf/TheTandon-Server/types"
    "golang.org/x/net/context"
    "errors"
)

const USER_INDEX = "test-user_data"
const USER_TYPE  = "user"

func IndexUser(eclient *elastic.Client, user types.User) error {
	ctx := context.Background()
    exists, err := eclient.IndexExists(USER_INDEX).Do(ctx)
    if err != nil {
        return err 
    }
    if !exists {
        return errors.New("Index does not exist")
    }
    _, err = eclient.Index().
        Index(USER_INDEX).
        Type(USER_TYPE).
        BodyJson(user).
        Do(ctx)
    if err != nil {
        return err 
    }
    return nil
}

func AddUserPrivilege(eclient *elastic.Client, user types.User, CourseName string, CourseID string, AccessLevel string) error {
    // First, we add Privilege to the user's privileges.
    user.Privileges = append(user.Privileges, types.Privilege{CourseName, CourseID, AccessLevel})

    // Finally, we index the updated user.
    return IndexUser(eclient, user)
}

func ModifyExistingUserPrivilege(eclient *elastic.Client, user types.User, CourseID string, NewAccessLevel string) error  {
    modified := false
    //  Update the user's access level.
    for index, copy := range user.Privileges {
        if (copy.CourseID == CourseID) {
            user.Privileges[index].AccessLevel = NewAccessLevel
            modified = true
            break
        }
    }
    //  If we can't find the privilege, throw an error
    if (!modified) {
        return errors.New("Privilege does not exist.")
    }
    // Index the updated user.
    return IndexUser(eclient, user)
}

func RemoveUserPrivilege(eclient *elastic.Client, user types.User, CourseID string) error {
    removed := false
    for index, copy := range user.Privileges {
        if (copy.CourseID == CourseID) {
            // Remove the privilege altogether
            user.Privileges = append(user.Privileges[:index], user.Privileges[index+1:]...)
            removed = true
        }
    }
    if (!removed) {
        return errors.New("Privilege does not exist.")
    }
    return IndexUser(eclient, user)
}
