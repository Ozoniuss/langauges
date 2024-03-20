# Installation and running comparison between languages


## C#

I've started this by first trying to understand the difference between the
.NET components, since you have runtime, core, sdk, framework. Wasn't sure what
needed to be installed to just run C#.

Once I read through the [glossary](https://learn.microsoft.com/en-us/dotnet/standard/glossary#net-5-and-later-versions)
I still found the namings unnecessarily confusing. But I went forward with 
installing .NET 8. 

```bash
sudo dnf install dotnet-sdk-8.0
```

I was not able to find simple steps to just create a c# file and compile it. All
the steps either show me how to use visual studio or using some stupid online
code execution tool. So fuck it, I settled for using `dotnet new console` and
got a shitload of files in my folder now.

I also added a C# extension that obviously wants my microsoft account and run
some automated setup. For fuck sake I just want to run hello world and have some
syntax highlight. Typical Microsoft.

Alright I was able to run the thing with `dotnet run`. Got my hello world in the
end.

Honestly, compared to other languages out there, this was much more complicated
than it had to be. The fact that having to [make a project seems simpler](https://stackoverflow.com/questions/48307508/how-can-i-run-individual-c-sharp-files-without-having-to-make-a-project) than
compiling the file directly is a bit of a turnoff. Languages like Python or Go
are obviously levels above in terms of simplicity, but I found even Java and
C++ (this one at least on Linux) to allow you to get started faster without any
fancy IDEs or Microsoft accounts. I understand that there is a reason behind all
those, but I'm just comparing how easy it is to have something quickly running.

Next, what of these project files do I even add on git??

I won't even complain about the fact that the C# extension did not work right
away. I don't even know how I got it working.
