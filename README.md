# CAMEL CASE

A program that converts a dash/dot/underscore/space separated string to camelCase or PascalCase.__
## Ex:
"foo-bar" becomes "fooBar"
Examples of how the module should be used; of course this would be in Go syntax

camelCase("foo-bar");
//=> "fooBar"

camelCase("foo_bar");
//=> "fooBar"

camelCase("Foo-Bar");
//=> "fooBar"

camelCase("Foo-Bar", { pascalCase: true });
//=> "FooBar"

camelCase("--foo.bar", { pascalCase: false });
//=> "fooBar"

camelCase("Foo-BAR", { uppercase: true });
//=> "fooBAR"

camelCase("fooBAR", { pascalCase: true, uppercase: true }));
//=> "FooBAR"

camelCase("foo bar");
//=> "fooBar"
