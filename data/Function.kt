package nameOfThePackage

import com.amazonaws.services.lambda.runtime.Context

class Function {

    class ExampleRequest {
        var hello: String = "azaza"

        constructor(hello: String) {
            this.hello = hello
        }

        constructor() {}
    }

    class ExampleResponse {
        var hello: String = "zazazaza"

        constructor(hello: String) {
            this.hello = hello
        }

        constructor() {}
    }

    fun handler(event: ExampleRequest, context: Context): ExampleResponse {
        return ExampleResponse(event.hello)
    }
}
