package bytegrind.oat.api;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.test.context.SpringBootTest;

import static org.junit.jupiter.api.Assertions.assertEquals;

@SpringBootTest(properties = "foo=bar")
public class OatApiApplicationTests {
  @Value("${foo}")
  String foo;

  @Test
  void contextLoads() {
    assertEquals("bar", foo);
  }
}
