package bytegrind.oat.api;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class HelloController {
  // http://localhost:8082/api/hello?name=World
  @GetMapping("/api/hello")
  public String sayHello(@RequestParam(defaultValue = "Guest") String name) {
    return "Hello, " + name + "!";
  }
}
