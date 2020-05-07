# What is it?

Acme-Twitter is a small set of command-line tools to interact with Twitter in a workflow that complements the Acme editor from Plan 9. Commands can be run in-line from the various tool outputs, such as retweet, reply, favorite, etc. Navigation is possible back into Twitter in the web browser using URL's. Since the output uses Markdown you can dump the content to a generator that will produce a prettier view of the content in HTML or even postscript/PDF for printing.

# How it works

Once you are all set up and authenticated with Twitter you can begin by bringing up your timeline in an Acme window using the acme-twitter or Ttl commands.

```
Ttl

__________________________
Tue May 5 14:41:35 (Tuser seeteegee) <https://twitter.com/seeteegee/status/1257742251654078471>

@rakyll This is really interesting. I wonder how to figure out a reasonable time threshold between the two events in order to serialize them?

(Trepl 1257742251654078471 '@seeteegee '):0 (Trt 1257742251654078471):0 (Tfav 1257742251654078471):0

__________________________
Tue May 5 14:43:25 (Tuser rakyll) <https://twitter.com/rakyll/status/1257742714856075264>

Most people think Spanner is about scaling data. Spanner's biggest feature is developer productivity in a service-oriented world. It also applies if you have a large number of clients working against the same data, no wonder it is extremely popular in gaming industry.

(Trepl 1257742714856075264 '@rakyll '):0 (Trt 1257742714856075264):0 (Tfav 1257742714856075264):0

__________________________
<Ttl -since 1257743153404182529
```

You can do a surprising number of things with this data inside Acme. First, you can easily get back to Twitter web application by right-clicking on the URL's. There are URL's for the individual tweets in the timeline as well as ones posted by the users themselves in their statuses.

If you select a line of underscores and right-click you can hop down tweet-by-tweet until you hit the end of the buffer. This also serves as a nice marker of where you left off in the stream. If you highlight a username you can similarly hop down through the tweets from that user.

When you get to the bottom you'll see a special "<Ttl -since" command. This is printed at the end to allow you to load additional tweets from that point into the buffer. Double click to the right of that line with the let mouse button to highligh the line and then middle-click in the selection to run the command. Now the buffer could look something like this and the process repeats itself.

```
Ttl

__________________________
Tue May 5 14:41:35 (Tuser seeteegee) <https://twitter.com/seeteegee/status/1257742251654078471>

@rakyll This is really interesting. I wonder how to figure out a reasonable time threshold between the two events in order to serialize them?

(Trepl 1257742251654078471 '@seeteegee '):0 (Trt 1257742251654078471):0 (Tfav 1257742251654078471):0

__________________________
Tue May 5 14:43:25 (Tuser rakyll) <https://twitter.com/rakyll/status/1257742714856075264>

Most people think Spanner is about scaling data. Spanner's biggest feature is developer productivity in a service-oriented worl
d. It also applies if you have a large number of clients working against the same data, no wonder it is extremely popular in gaming industry.

(Trepl 1257742714856075264 '@rakyll '):0 (Trt 1257742714856075264):0 (Tfav 1257742714856075264):0

__________________________
Tue May 5 14:33:34 (Tuser hillelogram) <https://twitter.com/hillelogram/status/1257740237423476737>

Someone plz test this, thx https://t.co/GAQhKNAWp5

(Trepl 1257740237423476737 '@hillelogram'):0 (Trt 1257740237423476737):0 (Tfav 1257740237423476737):0

__________________________
<Ttl -since 1257740237423476737
```

Note that there are three blocks at the end of each tweet partially in round brackets. The part inside the brackets are commands that you can run on the tweet: reply, retweet and favorite. To run then you will need to select the entire command. Luckily, there is a shortcut. If you double-click on one of the round brackets with the left mouse button it will select everything inside. Once the command is selected you can middle-click to run it. In the case of Trepl (reply) you will probably want to type your message inside the single quotes before selecting and running it.

There is another Tuser command that will bring up details about a particular user with further commands to see their timeline, follow, etc. though this is TBD.

To make a tweet use the Tweet command with your update as the first parameter (eg. ```Tweet 'Hello, World'```).

# Setup

Setup involves installing Go 1.13 or greater and running ```go install github.com/sirnewton01/acme-twitter```. Place the acme-twitter binary somewhere on your PATH where acme can run it. In that same location you can create symbolic links for Ttl, Trt, Trepl, Tfav, Tuser and Tweet pointing to the same acme-twitter executable. The acme-twitter command observes the name that it was invoked to decide which function to run.

# Authentication

Authentication is done for now using the same approach at the [t](https://github.com/sferik/t) command, which was the original inspiration for this tool. It uses the same authentication file at the moment with the same setup. You can follow the instructions there to get up and running. TBD: authentication should be done more independently.


