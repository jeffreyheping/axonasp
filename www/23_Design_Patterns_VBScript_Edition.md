# 23 Design Patterns Deep Dive — VBScript Edition

> VBScript supports `Class`, `Property Get/Let/Set`, and `Class_Initialize/Terminate`, but **does not support inheritance, interfaces, polymorphism, or overloading**. This code uses **composition over inheritance**, **uniform method naming conventions over interfaces**, and **conditional branching over polymorphism** to clearly express the core idea of each pattern.

---

## Chapter 1 Singleton

**Core Idea**: Only one instance exists globally.

**Example Explanation**: A script-level variable holds the sole instance, and creation is controlled via the `GetInstance` function. Two retrievals return the same object; modifying one affects the other.

```vbscript
' Script-level variable: global reference to the unique instance
Dim gInstance
Set gInstance = Nothing

Class Singleton
    Private m_Data

    ' Constructor: initialize default data
    Private Sub Class_Initialize
        m_Data = "I am the only instance"
    End Sub

    ' Read internal data
    Public Property Get Data
        Data = m_Data
    End Property

    ' Write internal data
    Public Property Let Data(value)
        m_Data = value
    End Property
End Class

' Global access point: create if instance does not exist, otherwise return existing one
' Returns: the unique instance of the Singleton class
Function GetInstance()
    If gInstance Is Nothing Then
        Set gInstance = New Singleton
    End If
    Set GetInstance = gInstance
End Function

' Demo: two retrievals yield the same object
Set s1 = GetInstance()
Set s2 = GetInstance()
s1.Data = "Modified"
Response.Write s2.Data   ' Modified (same object)
```

**VBScript Compromise Notes**:
- **No static variables**: VBScript classes do not support `Static` variables; a script-level (module-level) global variable `gInstance` must be used to hold the unique instance, breaking encapsulation.
- **Cannot prevent external New**: VBScript has no `Private` constructors or `Friend` access control, so external code can always `New Singleton` to bypass `GetInstance`, making a true singleton impossible to enforce.

---

## Chapter 2 Factory Method

**Core Idea**: Encapsulate the decision of "which object to create" inside a factory.

**Example Explanation**: Define two product classes, `Dog` and `Cat`. `AnimalFactory` decides which animal to create based on a passed type string. The caller does not need to know the concrete class names.

```vbscript
' Product class: Dog
Class Dog
    ' Make the dog bark
    Public Function Speak
        Response.Write "Woof"
    End Function
End Class

' Product class: Cat
Class Cat
    ' Make the cat meow
    Public Function Speak
        Response.Write "Meow"
    End Function
End Class

' Factory class: creates the corresponding animal object based on a type string
' animalType: "dog" or "cat"
' Returns: an instance of Dog or Cat
Class AnimalFactory
    Public Function CreateAnimal(animalType)
        Select Case LCase(animalType)
            Case "dog"
                Set CreateAnimal = New Dog
            Case "cat"
                Set CreateAnimal = New Cat
            Case Else
                Set CreateAnimal = Nothing
        End Select
    End Function
End Class

' Demo: create objects through the factory instead of direct New
Dim factory
Set factory = New AnimalFactory
Set myPet = factory.CreateAnimal("dog")
myPet.Speak   ' Woof
```

**VBScript Compromise Notes**:
- **No inheritance, no abstract methods**: The classic Factory Method relies on "abstract Creator + subclasses overriding FactoryMethod". VBScript has no inheritance, so the `Select Case` must be written inside the factory class. Adding a new product requires modifying the factory code, violating the Open/Closed Principle.
- **No interface constraints**: `Dog` and `Cat` have no `IAnimal` interface guaranteeing they both implement the `Speak` method. This relies entirely on developer discipline; passing a wrong type causes a runtime error.

---

## Chapter 3 Abstract Factory

**Core Idea**: Create families of related objects; switching the factory switches the entire style.

**Example Explanation**: `WinFactory` creates Windows-style buttons and checkboxes; `MacFactory` creates Mac-style buttons and checkboxes. Switching the factory switches the entire UI style without replacing components one by one.

```vbscript
' ===== Windows-style products =====
Class WinButton
    ' Draw a Windows-style button
    Public Function Paint
        Response.Write "Drawing Windows-style button"
    End Function
End Class

Class WinCheckbox
    ' Draw a Windows-style checkbox
    Public Function Paint
        Response.Write "Drawing Windows-style checkbox"
    End Function
End Class

' ===== Mac-style products =====
Class MacButton
    ' Draw a Mac-style button
    Public Function Paint
        Response.Write "Drawing Mac-style button"
    End Function
End Class

Class MacCheckbox
    ' Draw a Mac-style checkbox
    Public Function Paint
        Response.Write "Drawing Mac-style checkbox"
    End Function
End Class

' ===== Windows factory: creates a full set of Windows-style controls =====
Class WinFactory
    ' Create a Windows button
    Public Function CreateButton
        Set CreateButton = New WinButton
    End Function
    ' Create a Windows checkbox
    Public Function CreateCheckbox
        Set CreateCheckbox = New WinCheckbox
    End Function
End Class

' ===== Mac factory: creates a full set of Mac-style controls =====
Class MacFactory
    ' Create a Mac button
    Public Function CreateButton
        Set CreateButton = New MacButton
    End Function
    ' Create a Mac checkbox
    Public Function CreateCheckbox
        Set CreateCheckbox = New MacCheckbox
    End Function
End Class

' Demo: switch the factory, and the entire UI style changes
Dim uiFactory
Set uiFactory = New MacFactory   ' Change to WinFactory for Windows style
Dim btn, chk
Set btn = uiFactory.CreateButton
Set chk = uiFactory.CreateCheckbox
btn.Paint
chk.Paint
```

**VBScript Compromise Notes**:
- **No interfaces**: `WinFactory` and `MacFactory` share no common `IGUIFactory` interface, so the compiler cannot guarantee both implement `CreateButton`/`CreateCheckbox`. If a factory misses a method, the error occurs only at runtime.
- **No product constraints**: All Button and Checkbox classes rely solely on method name conventions; there is no `IButton`/`ICheckbox` interface ensuring consistency.

---

## Chapter 4 Builder

**Core Idea**: Construct complex objects step by step; the same construction process can produce different configurations.

**Example Explanation**: The `Director` instructs the `Builder` to assemble a `Computer` step by step (CPU -> RAM -> Disk). The same `Builder` can be directed to produce either a gaming PC or an office PC configuration.

```vbscript
' Product class: Computer
Class Computer
    Public CPU, RAM, Disk
    ' Print current configuration
    Public Function ShowConfig
        Response.Write "Config: " & CPU & " / " & RAM & " / " & Disk
    End Function
End Class

' Builder: assembles the Computer components step by step
Class ComputerBuilder
    Private m_Computer

    ' Constructor: create a blank product instance
    Private Sub Class_Initialize
        Set m_Computer = New Computer
    End Sub

    ' Install CPU
    Public Function BuildCPU(v)
        m_Computer.CPU = v
    End Function
    ' Install RAM
    Public Function BuildRAM(v)
        m_Computer.RAM = v
    End Function
    ' Install disk
    Public Function BuildDisk(v)
        m_Computer.Disk = v
    End Function
    ' Return the fully assembled product
    Public Function GetResult
        Set GetResult = m_Computer
    End Function
End Class

' Director: calls Builder steps in a fixed order, encapsulating different configuration schemes
Class Director
    ' Scheme 1: assemble a gaming PC
    Public Function ConstructGamingPC(builder)
        builder.BuildCPU "i9"
        builder.BuildRAM "32GB"
        builder.BuildDisk "2TB SSD"
    End Function
    ' Scheme 2: assemble an office PC
    Public Function ConstructOfficePC(builder)
        builder.BuildCPU "i5"
        builder.BuildRAM "16GB"
        builder.BuildDisk "512GB SSD"
    End Function
End Class

' Demo: Director instructs Builder to assemble
Dim builder, director, pc
Set builder = New ComputerBuilder
Set director = New Director
director.ConstructGamingPC builder
Set pc = builder.GetResult
pc.ShowConfig
```

**VBScript Compromise Notes**:
- **No interfaces**: The `builder` parameter received by `Director` has no `IBuilder` interface constraint. If the passed object lacks methods such as `BuildCPU`, the error occurs only at runtime.
- **No chaining**: Although VBScript `Function` can return `Me`, the syntax does not support chain calls like `builder.BuildCPU("i9").BuildRAM("32GB")` (`Set FunctionName = Me` works, but the caller must still receive return values line by line); only sequential calls are possible.

---

## Chapter 5 Prototype

**Core Idea**: Create new objects by copying existing ones.

**Example Explanation**: Create a resume template and duplicate it via the `Clone` method. Modifying the copy's skills array does not affect the original, demonstrating a deep copy.

```vbscript
' Resume class: contains name, age, and skills array
Class Resume
    Public Name, Age, Skills

    ' Deep-copy clone: create a new Resume and copy fields one by one
    ' Array is copied element by element to ensure modifying the copy does not affect the original
    ' Returns: a new Resume instance
    Public Function Clone
        Dim copy
        Set copy = New Resume
        copy.Name = Me.Name
        copy.Age = Me.Age
        Dim ub, i
        ub = UBound(Me.Skills)
        ReDim copy.Skills(ub)
        For i = 0 To ub
            copy.Skills(i) = Me.Skills(i)
        Next
        Set Clone = copy
    End Function
End Class

' Demo: after cloning, modifying the copy does not affect the original
Dim r1, r2
Set r1 = New Resume
r1.Name = "Zhang San"
r1.Age = 25
r1.Skills = Array("VBScript", "HTML")

Set r2 = r1.Clone
r2.Name = "Li Si"
r2.Skills(0) = "JavaScript"

Response.Write r1.Name & " " & r1.Skills(0)   ' Zhang San VBScript
Response.Write r2.Name & " " & r2.Skills(0)   ' Li Si JavaScript
```

**VBScript Compromise Notes**:
- **No built-in Clone**: VBScript has no `Clone()` method or serialization mechanism; fields must be copied manually one by one. The more fields there are, the more verbose the `Clone` method becomes, and it is easy to forget to add new fields to `Clone`.
- **No ICloneable interface**: There is no way to mandate that all classes implement `Clone`; it relies on developer discipline.

---

## Chapter 6 Proxy

**Core Idea**: Provide a surrogate for the real object to control access or delay loading.

**Example Explanation**: `ProxyImage` does not load the large image upon creation; only when `Display` is actually called does it create `RealImage` and load it. The second `Display` call directly reuses the already-loaded real object.

```vbscript
' Real object: loads and displays a large image
Class RealImage
    Private m_Filename

    ' Initialization: simulate a time-consuming large file load operation
    Public Function Init(filename)
        m_Filename = filename
        Response.Write "[Loading large image]" & filename
    End Function

    ' Display the image
    Public Function Display
        Response.Write "Displaying image: " & m_Filename
    End Function
End Class

' Proxy object: delays loading and controls access to RealImage
Class ProxyImage
    Private m_Filename
    Private m_RealImage   ' The proxied real object, initially Nothing

    ' Initialization: only record the filename, do not load
    Public Function Init(filename)
        m_Filename = filename
    End Function

    ' Display image: create the real object on first call, then reuse it
    Public Function Display
        If m_RealImage Is Nothing Then
            Set m_RealImage = New RealImage
            m_RealImage.Init m_Filename
        End If
        m_RealImage.Display
    End Function
End Class

' Demo: proxy creation does not load; Display triggers loading
Dim img
Set img = New ProxyImage
img.Init "photo.jpg"
Response.Write "Proxy created, real large image not yet loaded"
img.Display   ' Triggers real loading at this point
img.Display   ' No loading on second call
```

**VBScript Compromise Notes**:
- **No shared interface**: `ProxyImage` and `RealImage` share no `IImage` interface, so external code cannot transparently replace one with the other. The classic Proxy pattern requires the proxy and real object to implement the same interface, which VBScript cannot enforce.

---

## Chapter 7 Facade

**Core Idea**: Provide a simple, unified entry point for a complex subsystem.

**Example Explanation**: The boot process involves CPU Freeze -> Hard Drive Read -> Memory Load -> CPU Jump -> CPU Execute. `ComputerFacade` encapsulates these steps into a single `Start` method; the external caller only needs one call.

```vbscript
' Subsystem: CPU
Class CPU
    ' Freeze current state
    Public Function Freeze
        Response.Write "CPU Freeze"
    End Function
    ' Jump to specified address
    Public Function Jump(pos)
        Response.Write "CPU Jump to " & pos
    End Function
    ' Begin execution
    Public Function Execute
        Response.Write "CPU Execute"
    End Function
End Class

' Subsystem: Memory
Class Memory
    ' Load data into specified address
    Public Function Load(pos, data)
        Response.Write "Memory Load " & data & " to " & pos
    End Function
End Class

' Subsystem: Hard Drive
Class HardDrive
    ' Read data from specified sector
    ' Returns: simulated data block string
    Public Function Read(lba)
        Read = "Data block(" & lba & ")"
    End Function
End Class

' Facade class: encapsulates complex subsystem calls, exposing only Start externally
Class ComputerFacade
    Private m_CPU, m_Mem, m_HD

    ' Constructor: initialize all subsystems
    Private Sub Class_Initialize
        Set m_CPU = New CPU
        Set m_Mem = New Memory
        Set m_HD = New HardDrive
    End Sub

    ' One-key boot: internally calls each subsystem in order
    Public Function Start
        m_CPU.Freeze
        Dim bootData
        bootData = m_HD.Read(0)
        m_Mem.Load 0, bootData
        m_CPU.Jump 0
        m_CPU.Execute
    End Function
End Class

' Demo: external caller only needs Start, no need to understand internal details
Dim pc
Set pc = New ComputerFacade
pc.Start   ' Only Start is exposed externally
```

**VBScript Compromise Notes**: This pattern is relatively natural to implement in VBScript. The facade class simply composes and calls subsystems; it does not depend on inheritance or interfaces, so there are no significant compromises.

---

## Chapter 8 Adapter

**Core Idea**: Convert an incompatible interface into the target interface.

**Example Explanation**: `OldPrinter` only has an `OldPrint` method (accepting a string), while the new system expects a `Print` method (accepting a `Document` object). `PrinterAdapter` sits in between to convert: it extracts `doc.Content` and passes it to `OldPrint`.

```vbscript
' Old class: only has OldPrint method, accepts a string
Class OldPrinter
    ' Old interface: print string directly
    Public Function OldPrint(s)
        Response.Write "[Old Printer]" & s
    End Function
End Class

' Adapter: converts old interface to new interface
Class PrinterAdapter
    Private m_OldPrinter

    ' Inject the object to be adapted
    Public Function Init(oldPrinter)
        Set m_OldPrinter = oldPrinter
    End Function

    ' New interface: accepts Document object, extracts Content, and delegates to old interface
    Public Function Print(doc)
        m_OldPrinter.OldPrint doc.Content
    End Function
End Class

' New system data carrier
Class Document
    Public Content
End Class

' Demo: call old printer via new interface Print
Dim doc
Set doc = New Document
doc.Content = "Hello World"

Dim adapter
Set adapter = New PrinterAdapter
adapter.Init New OldPrinter
adapter.Print doc   ' Call old printer via new interface
```

**VBScript Compromise Notes**:
- **No target interface**: `PrinterAdapter` has no `IPrinter` interface to implement; the "new interface" is merely a `Print` method name convention. If there are multiple adapters, interface consistency cannot be guaranteed.

---

## Chapter 9 Bridge

**Core Idea**: Decouple abstraction from implementation so they can vary independently.

**Example Explanation**: `Circle` (abstraction layer) holds a reference to `Renderer` (implementation layer). The same `Circle` paired with `VectorRenderer` or `RasterRenderer` produces different rendering effects. Shapes and rendering engines can be extended independently.

```vbscript
' ===== Implementation layer: rendering engines (can be extended independently) =====

' Vector rendering engine
Class VectorRenderer
    ' Draw a circle using vector method
    Public Function RenderCircle(radius)
        Response.Write "Vector engine draws circle with radius " & radius
    End Function
End Class

' Raster rendering engine
Class RasterRenderer
    ' Draw a circle using raster method
    Public Function RenderCircle(radius)
        Response.Write "Raster engine draws circle with radius " & radius
    End Function
End Class

' ===== Abstraction layer: shape (holds a reference to implementation layer) =====

Class Circle
    Private m_Radius
    Private m_Renderer   ' Bridged rendering engine

    ' Initialization: pass radius and rendering engine
    Public Function Init(radius, renderer)
        m_Radius = radius
        Set m_Renderer = renderer
    End Function

    ' Draw: delegate to the held engine
    Public Function Draw
        m_Renderer.RenderCircle m_Radius
    End Function
End Class

' Demo: same shape, paired with different engines
Dim c1, c2
Set c1 = New Circle
c1.Init 5, New VectorRenderer
c1.Draw   ' Vector engine...

Set c2 = New Circle
c2.Init 5, New RasterRenderer
c2.Draw   ' Raster engine...
```

**VBScript Compromise Notes**:
- **No abstract classes**: The classic Bridge pattern requires the Abstraction (shape) to be an abstract class, with subclasses (circle, square) inheriting and extending it. VBScript has no inheritance; `Circle` is just a regular class and cannot form a Shape abstraction hierarchy.
- **No interface constraints**: Renderer has no `IRenderer` interface guaranteeing the existence of `RenderCircle`.

---

## Chapter 10 Composite

**Core Idea**: Treat individual objects and object compositions uniformly (tree structure).

**Example Explanation**: `Leaf` is a leaf node (employee), `Composite` is a composite node (department); both have the same `Operation` method. `Composite` internally recursively calls `Operation` on all child nodes, achieving "traverse leaf and branch with the same method".

```vbscript
' Leaf node: the terminal end of the tree structure
Class Leaf
    Public Name
    ' Display self info; indent controls indentation level
    Public Function Operation(indent)
        Response.Write indent & "Leaf: " & Name
    End Function
End Class

' Composite node: can contain child nodes (Leaf or Composite)
Class Composite
    Public Name
    Private m_Children()   ' Child node array
    Private m_Count        ' Current child count

    ' Constructor: initialize array
    Private Sub Class_Initialize
        m_Count = 0
        ReDim m_Children(10)
    End Sub

    ' Add child node (auto-expands when capacity is insufficient)
    Public Function Add(child)
        If m_Count > UBound(m_Children) Then
            ReDim Preserve m_Children(m_Count * 2)
        End If
        Set m_Children(m_Count) = child
        m_Count = m_Count + 1
    End Function

    ' Display self info and recursively call Operation on all child nodes
    Public Function Operation(indent)
        Response.Write indent & "Composite: " & Name
        Dim i
        For i = 0 To m_Count - 1
            m_Children(i).Operation indent & "  "
        Next
    End Function
End Class

' Demo: build a tree structure of HQ -> Branch -> Employees
Dim root, branch1, leaf1, leaf2, leaf3
Set root = New Composite
root.Name = "HQ"

Set branch1 = New Composite
branch1.Name = "Branch"

Set leaf1 = New Leaf
leaf1.Name = "Employee A"
Set leaf2 = New Leaf
leaf2.Name = "Employee B"
Set leaf3 = New Leaf
leaf3.Name = "Employee C"

branch1.Add leaf1
branch1.Add leaf2
root.Add branch1
root.Add leaf3

root.Operation ""   ' Uniformly traverse the entire tree
```

**VBScript Compromise Notes**:
- **No common base class**: The classic Composite pattern requires `Leaf` and `Composite` to inherit from the same `Component` base class. VBScript has no inheritance; the two are completely independent classes, relying solely on the `Operation` method name convention for "duck typing". The compiler cannot guarantee type safety.
- **No type safety**: The `Add` method accepts a `child` parameter with no type constraint; theoretically any object can be passed, and the error occurs only at runtime when `Operation` is called.

---

## Chapter 11 Decorator

**Core Idea**: Dynamically add new functionality to an object without modifying the original class.

**Example Explanation**: `SimpleCoffee` is the base object. `MilkDecorator` and `SugarDecorator` each hold a `coffee` reference. Passing `base` into `milk`, then `milk` into `sugar`, the final `sugar.Cost = base.Cost + 2 + 1`, stacking layer by layer.

```vbscript
' Base component: plain coffee
Class SimpleCoffee
    ' Return price
    Public Function Cost
        Cost = 10
    End Function
    ' Return description
    Public Function Description
        Description = "Plain coffee"
    End Function
End Class

' Decorator: milk (wraps a coffee object and adds price on top)
Class MilkDecorator
    Private m_Coffee   ' The wrapped inner object

    ' Inject the object to be decorated
    Public Function Init(coffee)
        Set m_Coffee = coffee
    End Function
    ' Price = inner object price + milk markup
    Public Function Cost
        Cost = m_Coffee.Cost + 2
    End Function
    ' Description = inner object description + milk
    Public Function Description
        Description = m_Coffee.Description & " + Milk"
    End Function
End Class

' Decorator: sugar (same structure as MilkDecorator)
Class SugarDecorator
    Private m_Coffee

    ' Inject the object to be decorated
    Public Function Init(coffee)
        Set m_Coffee = coffee
    End Function
    ' Price = inner object price + sugar markup
    Public Function Cost
        Cost = m_Coffee.Cost + 1
    End Function
    ' Description = inner object description + sugar
    Public Function Description
        Description = m_Coffee.Description & " + Sugar"
    End Function
End Class

' Demo: wrap layer by layer, dynamically stacking functionality
Dim base, milk, sugar
Set base = New SimpleCoffee
Response.Write base.Description & " = " & base.Cost & " yuan"

Set milk = New MilkDecorator
milk.Init base              ' milk wraps base

Set sugar = New SugarDecorator
sugar.Init milk             ' sugar wraps milk
Response.Write sugar.Description & " = " & sugar.Cost & " yuan"
```

**VBScript Compromise Notes**:
- **Cannot inherit base class**: The classic Decorator requires the Decorator to inherit from Component and hold a Component reference to achieve "type compatibility". VBScript has no inheritance; `MilkDecorator` and `SimpleCoffee` are completely different classes and cannot be substituted for each other. The caller must know whether it holds a decorator or the original object.
- **No transparency**: Ideally decorators are transparent to the caller, but in VBScript `Set milk = New MilkDecorator` and `Set coffee = New SimpleCoffee` are different types, and there is no way to declare a unified variable type.

---

## Chapter 12 Flyweight

**Core Idea**: Share fine-grained objects to reduce memory usage.

**Example Explanation**: A forest contains a large number of trees, but the "types" of trees (name + color) are only a few kinds. `TreeFactory` uses a `Dictionary` to cache `TreeType`; identical configurations are created only once, and multiple `Tree` instances share the same `TreeType`.

```vbscript
' Flyweight object: intrinsic attributes of a tree (name, color), shareable among multiple trees
Class TreeType
    Public Name, Color

    ' Draw the tree at specified coordinates
    Public Function Draw(x, y)
        Response.Write "Draw " & Color & Name & " at (" & x & "," & y & ")"
    End Function
End Class

' Flyweight factory: caches and reuses TreeType objects
Class TreeFactory
    Private m_Types   ' Dictionary: key -> TreeType

    ' Constructor: create dictionary
    Private Sub Class_Initialize
        Set m_Types = CreateObject("Scripting.Dictionary")
    End Sub

    ' Get or create TreeType: same parameters return the same object
    ' name: tree name, color: color
    ' Returns: shared TreeType instance
    Public Function GetTreeType(name, color)
        Dim key
        key = name & "|" & color
        If Not m_Types.Exists(key) Then
            Dim t
            Set t = New TreeType
            t.Name = name
            t.Color = color
            Set m_Types(key) = t
        End If
        Set GetTreeType = m_Types(key)
    End Function
End Class

' Demo: 3 trees share the same TreeType object
Dim factory, oakType, i
Set factory = New TreeFactory
Set oakType = factory.GetTreeType("Oak", "Green")

For i = 0 To 2
    oakType.Draw i, i * 2
Next
Response.Write "3 trees, but only 1 TreeType object actually exists"
```

**VBScript Compromise Notes**: This pattern is relatively natural to implement in VBScript. `Scripting.Dictionary` happens to provide the "cache objects by key" capability required by the Flyweight factory, aligning well with the pattern's needs. The only limitation is that Dictionary object storage and retrieval must explicitly use `Set`, making the syntax slightly verbose.

---

## Chapter 13 Strategy

**Core Idea**: Define a family of algorithms and make them interchangeable.

**Example Explanation**: Three pricing strategies (normal price, 20% off, 50% off). `ShoppingCart` holds the current strategy. At runtime, call `SetStrategy` to switch; `Checkout` automatically delegates to the current strategy for calculation. The caller does not need if-else logic.

```vbscript
' Strategy 1: normal price
Class NormalStrategy
    ' Return original price, no discount
    Public Function CalculatePrice(basePrice)
        CalculatePrice = basePrice
    End Function
End Class

' Strategy 2: 20% off
Class DiscountStrategy
    ' Apply 20% discount
    Public Function CalculatePrice(basePrice)
        CalculatePrice = basePrice * 0.8
    End Function
End Class

' Strategy 3: VIP 50% off
Class VipStrategy
    ' Apply 50% discount
    Public Function CalculatePrice(basePrice)
        CalculatePrice = basePrice * 0.5
    End Function
End Class

' Context: shopping cart, holds current strategy
Class ShoppingCart
    Private m_Strategy   ' Current pricing strategy
    Public Price         ' Original product price

    ' Switch pricing strategy
    Public Function SetStrategy(strategy)
        Set m_Strategy = strategy
    End Function

    ' Checkout: delegate calculation to current strategy
    Public Function Checkout
        Checkout = m_Strategy.CalculatePrice(Price)
    End Function
End Class

' Demo: switch strategies at runtime, same cart yields different prices
Dim cart
Set cart = New ShoppingCart
cart.Price = 100

cart.SetStrategy New NormalStrategy
Response.Write "Original: " & cart.Checkout    ' 100

cart.SetStrategy New DiscountStrategy
Response.Write "Discount: " & cart.Checkout  ' 80

cart.SetStrategy New VipStrategy
Response.Write "VIP: " & cart.Checkout   ' 50
```

**VBScript Compromise Notes**:
- **No interfaces**: The three strategies have no `IStrategy` interface guaranteeing they all have a `CalculatePrice` method. The `SetStrategy` parameter has no type constraint; passing a wrong object causes a runtime error.

---

## Chapter 14 Observer

**Core Idea**: When one object's state changes, automatically notify all objects observing it.

**Example Explanation**: `WeatherData` (subject) maintains a list of observers. When temperature changes, it calls `Notify`, iterating over all observers and calling their `Update` method. `PhoneDisplay` and `WindowDisplay` do not need to actively query; they passively receive notifications.

```vbscript
' Subject: weather data
Class WeatherData
    Private m_Observers()   ' Observer array
    Private m_Count         ' Observer count
    Public Temperature      ' Current temperature

    ' Constructor: initialize array
    Private Sub Class_Initialize
        m_Count = 0
        ReDim m_Observers(10)
    End Sub

    ' Register observer (auto-expands when capacity is insufficient)
    Public Function Register(observer)
        If m_Count > UBound(m_Observers) Then
            ReDim Preserve m_Observers(m_Count * 2)
        End If
        Set m_Observers(m_Count) = observer
        m_Count = m_Count + 1
    End Function

    ' Notify all observers: iterate and call each observer's Update
    Public Function Notify
        Dim i
        For i = 0 To m_Count - 1
            m_Observers(i).Update Me
        Next
    End Function

    ' Set temperature and trigger notification
    Public Function SetTemperature(t)
        Temperature = t
        Notify
    End Function
End Class

' Observer 1: phone display
Class PhoneDisplay
    ' Update display when notification is received
    Public Function Update(data)
        Response.Write "[Phone] Temperature: " & data.Temperature
    End Function
End Class

' Observer 2: window display
Class WindowDisplay
    ' Update display when notification is received
    Public Function Update(data)
        Response.Write "[Window] Temperature: " & data.Temperature
    End Function
End Class

' Demo: when temperature changes, both observers automatically receive notification
Dim weather
Set weather = New WeatherData
weather.Register New PhoneDisplay
weather.Register New WindowDisplay

weather.SetTemperature 25   ' Both observers receive notification simultaneously
weather.SetTemperature 30
```

**VBScript Compromise Notes**:
- **No interfaces**: `PhoneDisplay` and `WindowDisplay` have no `IObserver` interface guaranteeing they both have an `Update` method. The `Register` parameter has no type constraint.
- **No short-circuit evaluation**: VBScript's `And`/`Or` do not short-circuit. Although this example is not directly affected, conditional logic in notification code should be mindful of this limitation.

---

## Chapter 15 Template Method

**Core Idea**: Define the skeleton of an algorithm, delegating variable steps to external implementations.

**Example Explanation**: `DataMiner.Mine` is the template method, fixing the four steps "open file -> extract data -> analyze -> send report". "Extract data" is the variable step; switching is done by passing different `Extractor` objects (PDF/CSV).

```vbscript
' Template class: data miner, defines fixed algorithm skeleton
Class DataMiner
    ' Template method: fixed flow = open -> extract -> analyze -> send
    ' path: file path, extractor: extraction strategy object
    Public Function Mine(path, extractor)
        Dim file, rawData, analysis
        file = OpenFile(path)
        rawData = extractor.Extract(file)   ' Variable step delegated to extractor
        analysis = Analyze(rawData)
        SendReport analysis
    End Function

    ' Fixed step: simulate opening file
    Private Function OpenFile(path)
        OpenFile = "File content(" & path & ")"
    End Function
    ' Fixed step: simulate data analysis
    Private Function Analyze(data)
        Analyze = "Analysis result"
    End Function
    ' Fixed step: send report
    Private Function SendReport(r)
        Response.Write "Sending report: " & r
    End Function
End Class

' Variable step implementation 1: PDF extractor
Class PdfExtractor
    ' Extract data from PDF file
    Public Function Extract(file)
        Extract = "Extract from PDF: " & file
    End Function
End Class

' Variable step implementation 2: CSV extractor
Class CsvExtractor
    ' Extract data from CSV file
    Public Function Extract(file)
        Extract = "Extract from CSV: " & file
    End Function
End Class

' Demo: same skeleton, pass different extractors to handle different formats
Dim miner
Set miner = New DataMiner
miner.Mine "report.pdf", New PdfExtractor
miner.Mine "data.csv", New CsvExtractor
```

**VBScript Compromise Notes**:
- **No inheritance -> pattern essence changes**: The classic Template Method relies on "abstract base class defines template method + subclasses override variable steps". VBScript has no inheritance at all, so it can only be changed to **composition + delegation**: extract the variable step into an independent object passed in. This essentially shifts the pattern from "Template Method" to a variant of the "Strategy" pattern, altering the pattern's essence.
- **Cannot enforce skeleton immutability**: In the classic pattern, `Private` template methods + `Public` abstract methods ensure subclasses can only override variable steps. In VBScript, external code can directly call `OpenFile`/`Analyze`, making the skeleton unprotected.

---

## Chapter 16 Iterator

**Core Idea**: Provide a uniform way to traverse a collection without exposing its internal structure.

**Example Explanation**: `NameCollection` internally stores data in an array, but external code does not directly access the array. Through `CreateIterator`, an iterator is returned; `HasNext`/`NextItem` are used to access elements one by one, achieving "regardless of whether the internals are an array or a linked list, the traversal method is the same".

```vbscript
' Collection class: internally uses array storage, exposes only an iterator externally
Class NameCollection
    Private m_Items()
    Private m_Count

    ' Constructor: initialize array
    Private Sub Class_Initialize
        m_Count = 0
        ReDim m_Items(10)
    End Sub

    ' Add element (auto-expands when capacity is insufficient)
    Public Function Add(item)
        If m_Count > UBound(m_Items) Then
            ReDim Preserve m_Items(m_Count * 2)
        End If
        m_Items(m_Count) = item
        m_Count = m_Count + 1
    End Function

    ' Create iterator: pass out internal array and count
    ' Returns: ArrayIterator instance
    Public Function CreateIterator
        Dim it
        Set it = New ArrayIterator
        it.Init m_Items, m_Count
        Set CreateIterator = it
    End Function
End Class

' Iterator: encapsulates traversal logic
Class ArrayIterator
    Private m_Items   ' Data array copy
    Private m_Count   ' Total element count
    Private m_Index   ' Current cursor position

    ' Initialization: receive data and count, cursor resets to zero
    Public Function Init(items, count)
        m_Items = items
        m_Count = count
        m_Index = 0
    End Function

    ' Whether there is a next element
    ' Returns: True/False
    Public Function HasNext
        HasNext = (m_Index < m_Count)
    End Function

    ' Retrieve current element and advance cursor
    ' Returns: current element
    Public Function NextItem
        NextItem = m_Items(m_Index)
        m_Index = m_Index + 1
    End Function
End Class

' Demo: traverse using iterator without touching internal array
Dim names, it
Set names = New NameCollection
names.Add "Alice"
names.Add "Bob"
names.Add "Charlie"

Set it = names.CreateIterator
Do While it.HasNext
    Response.Write it.NextItem
Loop
```

**VBScript Compromise Notes**:
- **No For Each support**: VBScript's `For Each` can only iterate built-in collections and arrays; it does not support custom iterator protocols. The syntactic sugar `For Each item In names` is impossible; `HasNext`/`NextItem` must be called manually.
- **No IEnumerator interface**: The iterator has no interface constraint; `HasNext`/`NextItem` rely purely on method name conventions.

---

## Chapter 17 Chain of Responsibility

**Core Idea**: Pass a request along a chain until someone handles it.

**Example Explanation**: Three approvers form a chain: Team Leader (<= 100 yuan) -> Manager (<= 1000 yuan) -> Director (no upper limit). The request starts at the Team Leader; if it cannot be handled, it is passed to the next until someone handles it or the chain ends.

```vbscript
' Handler: a node on the chain
Class Handler
    Public Name             ' Role name (Team Leader / Manager / Director)
    Private m_Next          ' Next handler

    ' Set the next handler, forming a chain
    Public Function SetNext(h)
        Set m_Next = h
    End Function

    ' Handle request: handle if possible, otherwise pass to next
    ' amount: approval amount
    ' Returns: result string
    Public Function HandleRequest(amount)
        If CanHandle(amount) Then
            HandleRequest = Name & " handled " & amount & " yuan"
        ElseIf Not m_Next Is Nothing Then
            HandleRequest = m_Next.HandleRequest(amount)
        Else
            HandleRequest = "No one handled"
        End If
    End Function

    ' Determine whether current role can handle the amount (distinguished by Name)
    Private Function CanHandle(amount)
        Select Case Name
            Case "Team Leader"
                CanHandle = (amount <= 100)
            Case "Manager"
                CanHandle = (amount <= 1000)
            Case "Director"
                CanHandle = True
            Case Else
                CanHandle = False
        End Select
    End Function
End Class

' Demo: build approval chain Team Leader -> Manager -> Director
Dim leader, manager, director
Set leader = New Handler
leader.Name = "Team Leader"
Set manager = New Handler
manager.Name = "Manager"
Set director = New Handler
director.Name = "Director"

leader.SetNext manager
manager.SetNext director

Response.Write leader.HandleRequest(50)     ' Team Leader handles
Response.Write leader.HandleRequest(500)    ' Manager handles
Response.Write leader.HandleRequest(5000)   ' Director handles
```

**VBScript Compromise Notes**:
- **No inheritance -> simulating polymorphism with conditional branching**: The classic Chain of Responsibility uses "abstract Handler base class + subclasses override CanHandle". VBScript has no inheritance, so the permission judgments for all three roles must be crammed into a single `Select Case Name`, causing all logic to be squeezed into one class. Adding a new role requires modifying the internal `Handler` class code, violating the Open/Closed Principle.

---

## Chapter 18 Command

**Core Idea**: Encapsulate a request as an object, supporting undo and queuing.

**Example Explanation**: `Light` is the receiver; `LightOnCommand`/`LightOffCommand` are concrete commands (encapsulating operations on `Light`). `RemoteControl` holds the current command; pressing the button executes `Execute`, pressing undo calls `Undo`. Command objects decouple the "invoker" from the "receiver".

```vbscript
' Receiver: light (the object that actually performs actions)
Class Light
    ' Turn on
    Public Function TurnOn
        Response.Write "Light is on"
    End Function
    ' Turn off
    Public Function TurnOff
        Response.Write "Light is off"
    End Function
End Class

' Command: turn on (encapsulates "turn on" request as an object)
Class LightOnCommand
    Private m_Light

    ' Inject receiver
    Public Function Init(light)
        Set m_Light = light
    End Function
    ' Execute: turn on
    Public Function Execute
        m_Light.TurnOn
    End Function
    ' Undo: turn off (inverse of turn on)
    Public Function Undo
        m_Light.TurnOff
    End Function
End Class

' Command: turn off
Class LightOffCommand
    Private m_Light

    ' Inject receiver
    Public Function Init(light)
        Set m_Light = light
    End Function
    ' Execute: turn off
    Public Function Execute
        m_Light.TurnOff
    End Function
    ' Undo: turn on (inverse of turn off)
    Public Function Undo
        m_Light.TurnOn
    End Function
End Class

' Invoker: remote control (holds command, triggers execute and undo)
Class RemoteControl
    Private m_Command      ' Currently bound command
    Private m_LastCommand  ' Last executed command (for undo)

    ' Bind command
    Public Function SetCommand(cmd)
        Set m_Command = cmd
    End Function

    ' Press button: execute current command and record
    Public Function PressButton
        m_Command.Execute
        Set m_LastCommand = m_Command
    End Function

    ' Press undo: undo last executed command
    Public Function PressUndo
        If Not m_LastCommand Is Nothing Then
            m_LastCommand.Undo
        End If
    End Function
End Class

' Demo: commands can be undone after execution
Dim light, onCmd, offCmd, remote
Set light = New Light
Set onCmd = New LightOnCommand
onCmd.Init light
Set offCmd = New LightOffCommand
offCmd.Init light
Set remote = New RemoteControl

remote.SetCommand onCmd
remote.PressButton    ' Light is on
remote.SetCommand offCmd
remote.PressButton    ' Light is off
remote.PressUndo      ' Undo: Light is on
```

**VBScript Compromise Notes**:
- **No interfaces**: `LightOnCommand` and `LightOffCommand` have no `ICommand` interface guaranteeing they both have `Execute`/`Undo`. The `SetCommand` parameter has no type constraint.

---

## Chapter 19 State

**Core Idea**: When an object's internal state changes, its behavior changes accordingly.

**Example Explanation**: `Document` has three states: Draft -> Moderation -> Published. Each time `Publish` is called, the current state object decides the behavior: Draft state transitions to Moderation, Moderation transitions to Published, and Published indicates no further action needed. State transitions are driven by the state object itself.

```vbscript
' State: Draft
Class DraftState
    ' Publish action: draft can be published, transitions to moderation
    Public Function Publish(doc)
        Response.Write "Draft -> Moderation"
        doc.SetState "Moderation"
    End Function
End Class

' State: Moderation
Class ModerationState
    ' Publish action: moderation approved, transitions to published
    Public Function Publish(doc)
        Response.Write "Moderation -> Published"
        doc.SetState "Published"
    End Function
End Class

' State: Published
Class PublishedState
    ' Publish action: already published, no need to publish again
    Public Function Publish(doc)
        Response.Write "Already published, no need to publish again"
    End Function
End Class

' Context: document (holds current state, delegates behavior to state object)
Class Document
    Private m_State    ' Current state object
    Private m_States   ' Dictionary: state name -> state object

    ' Constructor: create all states, initial state is draft
    Private Sub Class_Initialize
        Set m_States = CreateObject("Scripting.Dictionary")
        Set m_States("Draft") = New DraftState
        Set m_States("Moderation") = New ModerationState
        Set m_States("Published") = New PublishedState
        Set m_State = m_States("Draft")
    End Sub

    ' Switch state
    Public Function SetState(name)
        Set m_State = m_States(name)
    End Function

    ' Publish: delegate to current state object for handling
    Public Function Publish
        m_State.Publish Me
    End Function
End Class

' Demo: consecutive publishes, state automatically transitions
Dim doc
Set doc = New Document
doc.Publish   ' Draft -> Moderation
doc.Publish   ' Moderation -> Published
doc.Publish   ' Already published...
```

**VBScript Compromise Notes**:
- **Using Dictionary instead of polymorphic dispatch**: The classic State pattern uses polymorphism to let different state subclasses automatically respond to the same method. VBScript has no polymorphism, so a `Dictionary` storing state objects + string key switching is used, essentially simulating polymorphic dispatch with a lookup table.
- **No interface constraints**: The three state classes have no `IState` interface guaranteeing they all have a `Publish` method; passing a wrong object causes a runtime error.

---

## Chapter 20 Mediator

**Core Idea**: Use a mediator object to encapsulate interactions between objects, reducing direct references.

**Example Explanation**: `User` objects do not communicate directly with each other; all communication goes through `ChatRoom`. `User` holds a reference to `ChatRoom`; when sending a message, it calls `ChatRoom.ShowMessage`. Adding a new user only requires injecting the same `ChatRoom`; the user does not need to know other users.

```vbscript
' Mediator: chat room (uniformly handles message display)
Class ChatRoom
    ' Display message: with timestamp and sender name
    ' user: sender, message: message content
    Public Function ShowMessage(user, message)
        Response.Write "[" & Now & "] " & user.Name & ": " & message
    End Function
End Class

' Colleague class: user (communicates via mediator, does not directly reference other users)
Class User
    Public Name
    Private m_ChatRoom   ' Held mediator reference

    ' Initialization: set name and mediator
    Public Function Init(name, chatRoom)
        Name = name
        Set m_ChatRoom = chatRoom
    End Function

    ' Send message: delegate to mediator
    Public Function SendMessage(message)
        m_ChatRoom.ShowMessage Me, message
    End Function
End Class

' Demo: two users communicate via chat room, no direct references to each other
Dim room, user1, user2
Set room = New ChatRoom
Set user1 = New User
user1.Init "Zhang San", room
Set user2 = New User
user2.Init "Li Si", room

user1.SendMessage "Hello"
user2.SendMessage "Hi"
```

**VBScript Compromise Notes**:
- **No interfaces**: `ChatRoom` has no `IMediator` interface, and `User` has no `IColleague` interface. The communication contract between colleague and mediator relies entirely on method name conventions. The classic pattern can enforce "all colleagues communicate through the mediator" via interfaces, which VBScript cannot enforce.

---

## Chapter 21 Visitor

**Core Idea**: Separate operations on elements of a structure into independent visitors.

**Example Explanation**: `Dot` and `Circle` are elements; `XMLExportVisitor` is a visitor. `ShapeCollection.Accept` iterates over all elements, uses `TypeName` to determine type, and calls the corresponding `VisitDot`/`VisitCircle`. Adding a new export format (e.g., JSON) only requires adding a new Visitor; element classes do not need to be modified.

```vbscript
' Element: Dot
Class Dot
    Public X, Y
End Class

' Element: Circle
Class Circle
    Public X, Y, Radius
End Class

' Visitor: XML export (has a corresponding Visit method for each element type)
Class XMLExportVisitor
    ' Visit Dot: output XML tag
    Public Function VisitDot(d)
        Response.Write "<dot x=""" & d.X & """ y=""" & d.Y & """/>"
    End Function
    ' Visit Circle: output XML tag
    Public Function VisitCircle(c)
        Response.Write "<circle x=""" & c.X & """ y=""" & c.Y & """ r=""" & c.Radius & """/>"
    End Function
End Class

' Object structure: shape collection, manages elements and accepts visitors
Class ShapeCollection
    Private m_Shapes()
    Private m_Count

    ' Constructor: initialize array
    Private Sub Class_Initialize
        m_Count = 0
        ReDim m_Shapes(10)
    End Sub

    ' Add shape (auto-expands when capacity is insufficient)
    Public Function Add(shape)
        If m_Count > UBound(m_Shapes) Then ReDim Preserve m_Shapes(m_Count * 2)
        Set m_Shapes(m_Count) = shape
        m_Count = m_Count + 1
    End Function

    ' Accept visitor: iterate elements, dispatch to corresponding Visit method by type
    ' Simulate double dispatch: VBScript has no overloading, use TypeName to determine type
    Public Function Accept(visitor)
        Dim i
        For i = 0 To m_Count - 1
            If TypeName(m_Shapes(i)) = "Dot" Then
                visitor.VisitDot m_Shapes(i)
            ElseIf TypeName(m_Shapes(i)) = "Circle" Then
                visitor.VisitCircle m_Shapes(i)
            End If
        Next
    End Function
End Class

' Demo: collection accepts XML visitor and automatically exports
Dim shapes, dot1, circle1
Set shapes = New ShapeCollection
Set dot1 = New Dot
dot1.X = 1
dot1.Y = 2
Set circle1 = New Circle
circle1.X = 3
circle1.Y = 4
circle1.Radius = 5
shapes.Add dot1
shapes.Add circle1

shapes.Accept New XMLExportVisitor
```

**VBScript Compromise Notes**:
- **No method overloading**: The classic Visitor uses `Visit(Dot)` and `Visit(Circle)` overloading to implement double dispatch. VBScript does not support overloading, so `VisitDot`/`VisitCircle` explicit naming must be used, and method names grow with element types.
- **Simulate double dispatch with TypeName()**: The classic pattern achieves double dispatch via "element.Accept(visitor) -> visitor.Visit(this)". VBScript has no polymorphism; `Accept` internally uses `TypeName()` string comparison to dispatch, which is runtime reflection rather than compile-time polymorphism, making it fragile and lacking type safety.

---

## Chapter 22 Memento

**Core Idea**: Capture and save an object's state to support rollback and recovery.

**Example Explanation**: `Editor` is the originator; the `Save` method creates an `EditorMemento` snapshot (saving `Content` and `CursorPos`). `History` is a stack; `Push` stores a snapshot, `Pop` retrieves it. `Restore` recovers state from the snapshot, implementing undo.

```vbscript
' Memento: saves a snapshot of Editor's state
Class EditorMemento
    Public Content, CursorPos
End Class

' Originator: editor (can save and restore state)
Class Editor
    Public Content, CursorPos

    ' Save current state to memento
    ' Returns: EditorMemento instance
    Public Function Save
        Dim m
        Set m = New EditorMemento
        m.Content = Content
        m.CursorPos = CursorPos
        Set Save = m
    End Function

    ' Restore state from memento
    Public Function Restore(memento)
        Content = memento.Content
        CursorPos = memento.CursorPos
    End Function
End Class

' Caretaker: history stack (stores mementos, supports undo)
Class History
    Private m_Stack()
    Private m_Count

    ' Constructor: initialize stack
    Private Sub Class_Initialize
        m_Count = 0
        ReDim m_Stack(10)
    End Sub

    ' Push: save memento (auto-expands when capacity is insufficient)
    Public Function Push(memento)
        If m_Count > UBound(m_Stack) Then ReDim Preserve m_Stack(m_Count * 2)
        Set m_Stack(m_Count) = memento
        m_Count = m_Count + 1
    End Function

    ' Pop: retrieve the most recent memento
    ' Returns: EditorMemento instance
    Public Function Pop
        m_Count = m_Count - 1
        Set Pop = m_Stack(m_Count)
    End Function
End Class

' Demo: edit -> save -> continue editing -> undo
Dim editor, history
Set editor = New Editor
Set history = New History

editor.Content = "Hello"
editor.CursorPos = 5
history.Push editor.Save       ' Save snapshot 1

editor.Content = "Hello World"
editor.CursorPos = 11

Response.Write "Current: " & editor.Content
editor.Restore history.Pop     ' Undo to snapshot 1
Response.Write "After undo: " & editor.Content
```

**VBScript Compromise Notes**: This pattern is relatively natural to implement in VBScript. The core of the Memento pattern is "state snapshot + recovery", which does not depend on inheritance or interfaces. VBScript's `Class` is sufficient to carry state data storage and recovery. The only deficiency is the inability to use access control (`Friend`/`Internal`) to restrict memento internal state access to the originator only; external code can directly access `EditorMemento`'s fields.

---

## Chapter 23 Interpreter

**Core Idea**: Define a grammar for a language and interpret sentences using an interpreter.

**Example Explanation**: Build an expression tree to check whether a text contains "Zhang San" or "Li Si". `TerminalExpression` is a leaf (matching a string), `OrExpression` is a composite (OR of two sub-expressions). `Interpret` recursively evaluates and returns True/False.

```vbscript
' Context: stores the input to be interpreted
Class Context
    Public Input
End Class

' Terminal expression: matches whether the input contains a specified keyword
Class TerminalExpression
    Private m_Data   ' Keyword to match

    ' Initialization: set keyword
    Public Function Init(data)
        m_Data = data
    End Function

    ' Interpret: determine whether the input contains the keyword
    ' Returns: True/False
    Public Function Interpret(context)
        Interpret = (InStr(context.Input, m_Data) > 0)
    End Function
End Class

' Non-terminal expression: OR operation (OR of two sub-expressions)
Class OrExpression
    Private m_Expr1, m_Expr2

    ' Initialization: inject two sub-expressions
    Public Function Init(expr1, expr2)
        Set m_Expr1 = expr1
        Set m_Expr2 = expr2
    End Function

    ' Interpret: True if either sub-expression is True
    Public Function Interpret(context)
        Interpret = m_Expr1.Interpret(context) Or m_Expr2.Interpret(context)
    End Function
End Class

' Non-terminal expression: AND operation (AND of two sub-expressions)
Class AndExpression
    Private m_Expr1, m_Expr2

    ' Initialization: inject two sub-expressions
    Public Function Init(expr1, expr2)
        Set m_Expr1 = expr1
        Set m_Expr2 = expr2
    End Function

    ' Interpret: True only if both sub-expressions are True
    Public Function Interpret(context)
        Interpret = m_Expr1.Interpret(context) And m_Expr2.Interpret(context)
    End Function
End Class

' Demo: build rule "Zhang San" or "Li Si", interpret different inputs
Dim exprZhang, exprLi, exprOr, ctx
Set exprZhang = New TerminalExpression
exprZhang.Init "Zhang San"
Set exprLi = New TerminalExpression
exprLi.Init "Li Si"
Set exprOr = New OrExpression
exprOr.Init exprZhang, exprLi

Set ctx = New Context
ctx.Input = "Participants: Zhang San, Wang Wu"
Response.Write "Match result: " & exprOr.Interpret(ctx)   ' True

ctx.Input = "Participants: Zhao Liu, Wang Wu"
Response.Write "Match result: " & exprOr.Interpret(ctx)   ' False
```

**VBScript Compromise Notes**:
- **No interfaces**: All expression classes have no `IExpression` interface guaranteeing they all have an `Interpret` method. The sub-expression references held by composite expressions have no type constraints.
- **No recursive type safety**: `OrExpression.Init` accepts `expr1`/`expr2` with no type annotation; theoretically any object can be passed, and the error occurs only at runtime when `Interpret` is called.

---

## Appendix: VBScript Syntax Quick Reference

| Feature | Syntax | Note |
|---------|--------|------|
| Define class | `Class ... End Class` | VBScript 5.0+ |
| Regular method | `Public Function Name ... End Function` | Always use Function; returns `Empty` when no value is returned |
| Constructor | `Private Sub Class_Initialize` | Must use Sub, language-enforced |
| Destructor | `Private Sub Class_Terminate` | Must use Sub, language-enforced |
| Property read | `Property Get Name ... End Property` | Returns value |
| Property write (value type) | `Property Let Name(v) ... End Property` | Writes scalar |
| Property write (object) | `Property Set Name(v) ... End Property` | Writes object reference |
| Create instance | `Set obj = New ClassName` | |
| Object assignment | `Set a = b` | Reference assignment must use Set |
| Dictionary store object | `Set dict(key) = obj` | Must use Set |
| Dictionary retrieve object | `Set obj = dict(key)` | Must use Set |
| Dictionary / Collection | `CreateObject("Scripting.Dictionary")` | Key-value container |
| Dynamic array | `ReDim Preserve arr(n)` | Resize while preserving data |
| Array upper bound | `UBound(arr)` | |
| Type check | `TypeName(obj)` | Returns class name string |
| Object null check | `obj Is Nothing` | |
| Short-circuit evaluation | **Not supported** | Both sides of `And`/`Or` are always evaluated |

---

## Appendix: 23 Patterns Overview

| # | Pattern | One-line Summary |
|---|---------|------------------|
| 1 | Singleton | Global unique instance |
| 2 | Factory Method | Conditional branching decides whom to create |
| 3 | Abstract Factory | Swap an entire family of related objects together |
| 4 | Builder | Assemble complex objects step by step |
| 5 | Prototype | Copy an existing object |
| 6 | Proxy | Surrogate controls access |
| 7 | Facade | Simple gateway to a complex system |
| 8 | Adapter | Interface converter |
| 9 | Bridge | Abstraction and implementation vary independently |
| 10 | Composite | Treat individual and composite uniformly |
| 11 | Decorator | Wrap layer by layer to add functionality |
| 12 | Flyweight | Share objects to save memory |
| 13 | Strategy | Switch algorithms at runtime |
| 14 | Observer | State changes, notify everyone |
| 15 | Template Method | Fixed skeleton, variable steps delegated |
| 16 | Iterator | Unified traversal interface |
| 17 | Chain of Responsibility | Request passes along chain until handled |
| 18 | Command | Encapsulate request as object, support undo |
| 19 | State | State changes, behavior follows |
| 20 | Mediator | Everyone talks through mediator instead of directly |
| 21 | Visitor | Separate operations from data structure |
| 22 | Memento | Save snapshot, support rollback |
| 23 | Interpreter | Define grammar and interpret execution |
