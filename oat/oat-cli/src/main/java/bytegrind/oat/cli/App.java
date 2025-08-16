package bytegrind.oat.cli;

import bytegrind.oat.infra.Test;

public class App {
    public static void main(String[] args) {
        final Test t = new Test();
        System.out.println("App " + t.test());
    }
}
