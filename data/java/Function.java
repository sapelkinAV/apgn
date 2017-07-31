package nameOfThePackage;

import com.amazonaws.services.lambda.runtime.Context;

public class Function {

    public static class FunctionRequest {
        String hello;

        public String getHello() {
            return hello;
        }

        public void setHello(String hello) {
            this.hello = hello;
        }

        public FunctionRequest(String hello) {
            this.hello = hello;
        }

        public FunctionRequest() {
        }
    }

    public static class FunctionResponse {
        String hello;

        public String getHello() {
            return hello;
        }

        public void setHello(String hello) {
            this.hello = hello;
        }

        public FunctionResponse(String hello) {
            this.hello = hello;
        }

        public FunctionResponse() {
        }
    }

    public FunctionResponse handler(FunctionRequest event, Context context) {
        return new FunctionResponse(event.getHello());
    }
}