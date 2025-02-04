import androidx.compose.ui.window.ComposeUIViewController
import ru.seventhrombus.app.App
import platform.UIKit.UIViewController

fun MainViewController(): UIViewController = ComposeUIViewController { App() }
