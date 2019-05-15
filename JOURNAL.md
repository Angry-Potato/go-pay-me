## 9th May

With no time for any real work tonight, I decided to read the brief a bunch of times and do a little research; I want to be clear about what's expected, and start exploring my options.

I find the brief unusual. The requirement of an API design prior to implementation is unexpected, and the language of the design requirement is perhaps purposefully ambiguous.

I like the sound of formally designing the API prior to implementation - it's like TDD at the service level.

The most confusing part of the brief is:

> _"Provide a design for part of a payments API."_

What does "part" of an API mean? I will assume it's reference to the fact that this particular API will be part of a group of services that are collectively responsible for payments, e.g. this payments API will expose payments resources whilst another service enforces processing of those payments.

I believe the vague language of the design requirement is deliberate. The only design format stipulation being that the document is PDF tells me that a simple OpenAPI/Swagger JSON file won't cut it.

I believe, to satisfy the design aspect of the brief, I will supply an OpenAPI driven design specification, and perhaps a simple wireframe demonstrating a possible client GUI, in a single PDF.

The implementation aspect should be simple. There are numerous Golang REST API web frameworks to choose from, I imagine I won't need to stray far past any auto-generation they may offer. Any development I undertake will be done using TDD. There should be a detailed README instructing users how they might build, run, and test the API, and where they might find it running in the wild.

I'd like to find a UI language / framework I'm unfamiliar with, and throw a small GUI up to consume from the API.

These might be some best laid plans, we shall see.

## 10th May

I remembered Google's Protobuf is a thing today. Could be interesting to use this to define / design an API prior to implementing it. Maybe look at other service contract systems like Pact.

After looking in more detail at Protobuf, and Pact, it still looks like a swagger / openapi API design document is the best choice for this project because swagger / openapi is specifically about documenting APIs for human readers.

It looks like it's possible to generate a swagger JSON file from scratch, then convert this file to AsciiDoc, and convert _that_ file to PDF.

Spent some time testing, writing, and documenting a simple script tool that will convert a swagger JSON file to a PDF. The tool is dockerised for convenience. I just need to write a swagger file based on the given example payments list, and the PDF is taken care of!

## 11th May

Going to write the swagger file this morning, and look at possible GUI design tools.

## 12th May

Should get the swagger file finished this morning, should have finished yesterday but it took longer than expected to learn the structure.

Pretty happy with the finished swagger file, it meets all of the requirements in the brief. The bulk update and bulk delete are not specified in the brief but I decided to keep them anyway to demonstrate my understanding of a RESTful API.

I noticed that the API design spec in the brief talks about more than just the REST interface. This leads me to believe I should add the tech choices to the design file.

I've looked at many Golang web frameworks, and landed on [go-json-rest](https://github.com/ant0ine/go-json-rest) backed by [gorm](https://github.com/jinzhu/gorm) connecting to a [postgres](https://www.postgresql.org/) database. The reasoning behind these choices is largely to do with ease of use (though I've never used them before, the documentation is high quality) which hopefully leads to speedy development. In addition, I was looking for libraries / frameworks that were very popular (github stars!) as this is usually a sign of a good library / framework, and it also means issues are likely to be more google-able.

Given that my favourite part of development is automation, I see it as important that I get the implementation's pipeline up before anything else. I've implemented the default [go-json-rest](https://github.com/ant0ine/go-json-rest) server, and written a single test, I'll now hook this up to [Travis-CI](https://travis-ci.org/). I chose to automate using Travis, even though I have no experience with it, because I saw that Form3 use Travis for at least some of their projects.

I have built the Travis pipeline! It builds and tests the docker image, then deploys to [heroku](https://go-pay-me.herokuapp.com/). Suppose I should do the feature work now..

## 14th May

Yesterday I finished implementing the basic payments RESTful interactions. The implementation took a bit longer than expected as I failed to account for library learning time. I'm happy with the automation, and design doc generation that I've set up. Less happy with the test suite in the implementation, there are a couple of race conditions. I'm contemplating writing a small acceptance test suite that will run in series to verify correctness of the API. At the very least, the existing suite needs to be refactored as it's pretty WET.

I need to update READMEs with build/test/run instructions, and add the sub-resources to payments, and that should be it!

I moved the swagger2pdf tool out to a separate repo, it's really a separate tool.

Went a bit nuts with the emoji's in the readme. Hopefully it's received as playful rather than irritating.

Began separating out unit / integration / acceptance tests because I want to be able to verify complete correctness of the app. It's hard to assert your `DELETE` operation deletes all resources when another test has just `POST`ed some new ones. To do this, the acceptance suite will run in isolation, in series, after the asynchronous (unit & integration) tests. This could be made quicker by upping two instances of the app to test async & sync suites in parallel but I feel that's overkill for this little project, and the logging output would be confusing at best.

## 15th May

I am coming to the end of the project now, and I have some regrets.

I should have started by planning the work, and tracking it in trello / github issues. This would have demonstrated responsible workflow management skills, of which I guess I have none!

I should have used testing frameworks like ginkgo & gomega to give me niceties like setup & teardown. What stopped me from using them was past experience - I found them to be tricky to set up - but this should not have stopped me, after all I am a far better golanger than I was back then.
