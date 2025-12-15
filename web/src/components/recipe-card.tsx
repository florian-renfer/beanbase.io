import { Badge } from "@/components/ui/badge";
import {
  Card,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Separator } from "@/components/ui/separator";
import { CoffeeIcon } from "@hugeicons/core-free-icons";
import { HugeiconsIcon } from "@hugeicons/react";

function RecipeCard() {
  return (
    <Card className="relative w-full max-w-sm overflow-hidden pt-0">
      <div className="absolute inset-0 z-30 aspect-video bg-primary opacity-50 mix-blend-color" />
      <img
        src="https://images.unsplash.com/photo-1604076850742-4c7221f3101b?q=80&w=1887&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"
        alt="Photo by mymind on Unsplash"
        title="Photo by mymind on Unsplash"
        className="relative z-20 aspect-video w-full object-cover brightness-60 grayscale"
      />
      <CardHeader>
        <div className="inline-flex w-full">
          <CardTitle>Espresso - Amabile</CardTitle>
          <Badge variant="default" className="ml-auto">
            Daily
          </Badge>
        </div>

        <CardDescription>
          <p>A delicious espresso base for all kind of drinks and recipes.</p>
        </CardDescription>
      </CardHeader>
      <CardFooter>
        <div className="flex h-5 items-center space-x-4 text-sm">
          <HugeiconsIcon icon={CoffeeIcon} />
          <Separator orientation="vertical" />
          <div>Parrotcaffe</div>
          <Separator orientation="vertical" />
          <div>Colombia</div>
        </div>
      </CardFooter>
    </Card>
  );
}

export { RecipeCard };
