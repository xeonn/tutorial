package my.onn.sb.kc.demo;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;

@Controller
public class CustomerController {

    @Autowired
    private CustomerDao customerDao;

    @GetMapping(path = "/")
    public String index() {
        return "external";
    }

    @GetMapping(path = "/customers")
    public String customers(Model model) {
        addCustomers();
        model.addAttribute("customers", customerDao.findAll());
        return "customers";
    }

    private void addCustomers() {
        customerDao.save(Customer.builder()
                .address("111 foo blvd")
                .name("Foo Industries")
                .serviceRendered("Important services")
                .build());
        customerDao.save(Customer.builder()
                .address("222 bar blvd")
                .name("Bar Industries")
                .serviceRendered("Important services")
                .build());
        customerDao.save(Customer.builder()
                .address("333 main street")
                .name("Big LLC")
                .serviceRendered("Important services")
                .build());
    }
}
