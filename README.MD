###URL Grabber Tool

It's simply allow you to GET from URL lots of files. It checks on 404 page status for filtering garbage output, but not anything else like 502 or so.

Also you can find BINNARYS for every popular OS/Architecture.

###Exapmle:

Like this:
https://example.com/dir/dir/*IDWITH_NUMBERS*/somthing/image.jpg

Or this:
https://example.com/dir/dir/DCIM*IDWITH_NUMBERS*.raw

Or even this:
https://example.com/dir/dir/*IDWITH_NUMBERS*.zip

###Config:

You must set ENV variables for:

*URLHEAD* - thas is beggining part of URL adress as ("https://example.com/dir/dir/") or ("https://example.com/dir/dir/DCIM"), so before the numbers begging.

*URLEND* - last part of URL adress after numbers as ("/somthing/image.jpg") or (".raw") - preaty much every constant in URL after numbers itteration.

*HEADITTER* - beggining number of file or directory.

*ENDITTER* - last number in file or directory.

*FILEFORMAT* - it's simply output: .jpg, .raw, .zip (you also neet to provide DOT, cuz reasons)


All output will be in to ./damp floder.
