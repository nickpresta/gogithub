# gogithub

An incomplete Go implementation of the GitHub API v3 (http://developer.github.com/v3/).

## Getting the Code

    go get github.com/NickPresta/gogithub

## Usage

Import this library in your code to use it:

    import (
        "github.com/NickPresta/gogithub"
    )

Fetch user details:

    // Anonymous querying
    credentials = map[string]string{}
    client, err := gogithub.Client(credentials)

    user, err := client.GetUser("NickPresta")
    // Show my Avatar URL
    fmt.Println(user.AvatarUrl)

See `examples/main.go` for an example.

## Documentation

View the documentation on
[GoPkgDoc](http://go.pkgdoc.org/github.com/NickPresta/gogithub).

## Tests

There are no tests at the moment.

## Changelog

See `CHANGELOG.md` for details.

## License

This is released under the
[MIT license](http://www.opensource.org/licenses/mit-license.php).
See `LICENSE.md` for details.
