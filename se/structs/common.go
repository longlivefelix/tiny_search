/*
@Author: Felix <https://github.com/longlivefelix>
@Date:   2020-04-21 18:56
*/
package structs

type SearchErr struct{
	level string
	msg string
}
func (err SearchErr)Error() string{
	return err.level + ": " +err.msg
}
